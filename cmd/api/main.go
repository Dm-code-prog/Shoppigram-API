package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/shoppigram-com/marketplace-api/packages/cloudwatchcollector"
	"github.com/shoppigram-com/marketplace-api/packages/cors"
	"github.com/shoppigram-com/marketplace-api/packages/health"
	"github.com/shoppigram-com/marketplace-api/packages/httpmetrics"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/shoppigram-com/marketplace-api/internal/app"

	"github.com/shoppigram-com/marketplace-api/internal/webhooks"

	"github.com/Netflix/go-env"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/run"
	"github.com/shoppigram-com/marketplace-api/internal/admin"
	"github.com/shoppigram-com/marketplace-api/internal/auth"
	"github.com/shoppigram-com/marketplace-api/internal/notifications"
	"go.uber.org/zap"
)

func main() {
	var config Environment
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		fmt.Println("failed to load environment variables", logger.SilentError(err))
		os.Exit(-1)
	}

	log := logger.New(config.Zap.LogLevel)

	cloudwatchcollector.Init(config.AWS.Cloudwatch.Namespace)
	defer cloudwatchcollector.Shutdown()

	pgConf, err := pgxpool.ParseConfig(config.Postgres.DSN)
	if err != nil {
		log.Fatal("failed to parse postgres dsn", logger.SilentError(err))
	}

	pgConf.MaxConns = int32(config.Postgres.MaxConns)
	pgConf.MinConns = int32(config.Postgres.MinConns)

	ctx := context.Background()
	db, err := pgxpool.NewWithConfig(ctx, pgConf)
	if err != nil {
		log.Fatal("failed to connect to database", logger.SilentError(err))
	}

	defer db.Close()
	log.Debug("connected to database")

	var g run.Group
	g.Add(run.SignalHandler(ctx, os.Interrupt, os.Kill, syscall.SIGTERM))

	var r = chi.NewRouter()
	var httpServer = http.Server{
		Addr:    ":" + config.HTTP.Port,
		Handler: r,
	}

	r.Use(
		middleware.Timeout(10*time.Second),
		middleware.Recoverer,
		middleware.Compress(5, "application/json"),
		cors.MakeCORSMiddleware(
			[]string{
				"https://web-app.shoppigram.com",
				"https://admin.shoppigram.com",
				"https://dev-app.shoppigram.com",
				"http://localhost:5173",
			},
			nil,
		),
		httpmetrics.MakeObservabilityMiddleware,
		middleware.Throttle(100),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "the path you requested does not exist"})
	})

	s3Instance := s3.New(
		session.Must(session.NewSession(&aws.Config{
			Region: aws.String("fra1"),
			Credentials: credentials.NewStaticCredentials(
				config.AWS.S3.Key,
				config.AWS.S3.Secret,
				"",
			),
			Endpoint:         aws.String(config.AWS.S3.Endpoint),
			S3ForcePathStyle: aws.Bool(false),
		})))

	////////////////////////////////////// TELEGRAM USERS //////////////////////////////////////
	authMw := auth.MakeAuthMiddleware(config.Bot.Token)
	tgUsersRepo := auth.NewPg(db)
	tgUsersService := auth.NewServiceWithObservability(
		auth.New(tgUsersRepo),
		log.With(zap.String("service", "telegram_users")),
	)
	authHandler := auth.MakeHandler(tgUsersService, authMw)

	////////////////////////////////////// MARKETPLACES //////////////////////////////////////
	marketplacesRepo := app.NewPg(db)
	marketplacesService := app.NewServiceWithObservability(
		app.New(marketplacesRepo, config.Cache.MaxSize),
		log.With(zap.String("service", "marketplaces")),
	)
	shopHandler := app.MakeShopHandler(marketplacesService)
	ordersHandler := app.MakeOrdersHandler(marketplacesService, authMw)

	////////////////////////////////////// NOTIFICATIONS //////////////////////////////////////
	notificationsRepo := notifications.NewPg(
		db,
		config.NewOrderNotifications.BatchSize,
		config.NewMarketplaceNotifications.BatchSize,
		config.VerifiedMarketplaceNotifications.BatchSize,
	)

	notificationsService := notifications.New(
		notificationsRepo,
		log.With(zap.String("service", "notifications")),
		time.Duration(config.NewOrderNotifications.TimeoutSec)*time.Second,
		time.Duration(config.NewMarketplaceNotifications.TimeoutSec)*time.Second,
		time.Duration(config.VerifiedMarketplaceNotifications.TimeoutSec)*time.Second,
		config.Bot.Token,
		config.Bot.Name,
	)

	////////////////////////////////////// RUN NOTIFICATION JOBS //////////////////////////////////////
	if config.NewOrderNotifications.IsEnabled {
		g.Add(notificationsService.RunOrdersNotifier, func(err error) {
			_ = notificationsService.Shutdown()
		})
	} else {
		log.Warn("new order notifications job is disabled")
	}

	if config.NewMarketplaceNotifications.IsEnabled {
		g.Add(notificationsService.RunNewMarketplaceNotifier, func(err error) {
			_ = notificationsService.Shutdown()
		})
	} else {
		log.Warn("new marketplace notifications job is disabled")
	}

	if config.VerifiedMarketplaceNotifications.IsEnabled {
		g.Add(notificationsService.RunVerifiedMarketplaceNotifier, func(err error) {
			_ = notificationsService.Shutdown()
		})
	} else {
		log.Warn("verified marketplace notifications job is disabled")
	}

	////////////////////////////////////// ADMINS /////////////////////////////////////

	adminsRepo := admin.NewPg(db)
	adminsService := admin.NewServiceWithObservability(
		admin.New(
			adminsRepo,
			&notificationsAdminAdapter{
				notifier: notificationsService,
			},
			s3Instance,
			config.Bot.Name,
			config.AWS.S3.Bucket,
		),
		log.With(zap.String("service", "admins")),
	)
	adminsHandler := admin.MakeHandler(adminsService, authMw)
	adminsHandlerV2 := admin.MakeHandlerV2(adminsService, authMw)

	////////////////////////////////////// WEBHOOKS //////////////////////////////////////
	tgRepo := webhooks.Repository(webhooks.NewPg(db))
	webhookService := webhooks.NewTelegram(
		tgRepo,
		&notificationsWebhooksAdapter{notifier: notificationsService},
		log.With(zap.String("service", "webhooks")),
		config.Bot.ID,
		config.Bot.Name,
	)
	webhooksHandler := webhooks.MakeTelegramHandler(
		webhookService,
		log.With(zap.String("service", "webhooks_server")),
		config.TelegramWebhooks.SecretToken)

	webhooksRepo := webhooks.NewPg(db)
	maxCloudPaymentsTransactionDuration, _ := time.ParseDuration(config.CloudPayments.MaxTransactionDuration)
	cloudPaymentsWebhookService := webhooks.NewCloudPayments(
		webhooksRepo,
		log.With(zap.String("service", "webhooks_server")),
		maxCloudPaymentsTransactionDuration,
	)
	cloudPaymentsWebhookHandler := webhooks.MakeCloudPaymentsHandlers(
		cloudPaymentsWebhookService,
		log.With(zap.String("service", "webhooks_server")),
		config.CloudPayments.Login,
		config.CloudPayments.Password,
	)

	////////////////////////////////////// HTTP SERVER V1 //////////////////////////////////////
	r.Mount("/api/v1/public/products", shopHandler)
	r.Mount("/api/v1/public/auth", authHandler)
	r.Mount("/api/v1/public/orders", ordersHandler)
	r.Mount("/api/v1/private/marketplaces", adminsHandler)

	r.Mount("/api/v1/telegram/webhooks", webhooksHandler)
	r.Mount("/api/v1/cloud-payments/webhooks", cloudPaymentsWebhookHandler)

	////////////////////////////////////// RUN HTTP SERVER V2 //////////////////////////////////////

	r.Mount("/api/v2/app/shops", shopHandler)
	r.Mount("/api/v2/app/orders", ordersHandler)
	r.Mount("/api/v2/auth", authHandler)
	r.Mount("/api/v2/admin", adminsHandlerV2)

	g.Add(func() error {
		log.Info("starting HTTP server", zap.String("port", config.HTTP.Port))
		return httpServer.ListenAndServe()
	}, func(err error) {
		if err != nil {
			log.Error("HTTP server exited with error", logger.SilentError(err))
		}
		err = httpServer.Shutdown(ctx)
		if err != nil {
			log.Error("failed to shutdown HTTP server", logger.SilentError(err))
		}
	})

	healthR := chi.NewRouter()
	healthR.Mount("/health", health.NewHandler())
	healthServer := &http.Server{
		Addr:    ":7777",
		Handler: healthR,
	}
	g.Add(healthServer.ListenAndServe, func(err error) {
		if err != nil {
			log.Error("health server exited with error", logger.SilentError(err))
		}
		err = healthServer.Shutdown(ctx)
		if err != nil {
			log.Error("failed to shutdown health server", logger.SilentError(err))
		}
	})

	if err := g.Run(); err != nil {
		log.Info("api exited with message:", logger.SilentError(err))
		return
	}
}
