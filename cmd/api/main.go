package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shoppigram-com/marketplace-api/internal/marketplaces"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/shoppigram-com/marketplace-api/internal/webhooks"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shoppigram-com/marketplace-api/internal/admins"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"github.com/shoppigram-com/marketplace-api/internal/notifications"
	"go.uber.org/zap/zapcore"

	"github.com/Netflix/go-env"
	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oklog/run"
	"github.com/shoppigram-com/marketplace-api/internal/httputils"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
)

func main() {
	var (
		logLevel  zapcore.Level
		zapConfig zap.Config
	)

	ctx := context.Background()

	var config Environment
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		fmt.Println("failed to load environment variables", logging.SilentError(err))
		return
	}

	switch config.Zap.LogLevel {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "WARN", "WARNING":
		logLevel = zapcore.WarnLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	if config.Zap.LogLevel == "DEBUG" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Level.SetLevel(logLevel)

	log, _ := zapConfig.Build(zap.AddStacktrace(zapcore.PanicLevel))
	defer log.Sync()

	db, err := pgxpool.New(ctx, config.Postgres.DSN)
	if err != nil {
		log.Fatal("failed to connect to database", logging.SilentError(err))
		return
	}
	defer db.Close()
	log.Debug("connected to database")

	db.Config().MinConns = 5
	db.Config().MaxConns = 25

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
		httputils.CORSMiddleware,
		middleware.Throttle(500),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "the path you requested does not exist"})
	})

	productsCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,         // number of keys to track frequency of (10M).
		MaxCost:     200_000_000, // maximum cost of productsCache (200 MB).
		BufferItems: 64,          // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("failed to create productsCache", logging.SilentError(err))
		return
	}

	authMw := telegramusers.MakeAuthMiddleware(log.With(zap.String("service", "users")), config.Bot.Token)

	tgUsersRepo := telegramusers.NewPg(db, config.Encryption.Key)
	tgUsersService := telegramusers.New(tgUsersRepo, log.With(zap.String("service", "users")))
	tgUsersHandler := telegramusers.MakeHandler(tgUsersService, authMw)

	marketplacesRepo := marketplaces.NewPg(db)
	marketplacesService := marketplaces.New(marketplacesRepo, log.With(zap.String("service", "marketplaces")), productsCache)
	productsHandler := marketplaces.MakeProductsHandler(marketplacesService)
	ordersHandler := marketplaces.MakeOrdersHandler(marketplacesService, authMw)

	notificationsRepo := notifications.NewPg(
		db,
		config.NewOrderNotifications.BatchSize,
		config.NewMarketplaceNotifications.BatchSize,
		config.VerifiedMarketplaceNotifications.BatchSize,
	)
	notificationsService := notifications.New(
		notificationsRepo,
		log.With(zap.String("service", "notifications")),
		time.Duration(config.NewOrderNotifications.Timeout)*time.Second,
		time.Duration(config.NewMarketplaceNotifications.Timeout)*time.Second,
		time.Duration(config.VerifiedMarketplaceNotifications.Timeout)*time.Second,
		config.Bot.Token,
	)

	if config.NewOrderNotifications.IsEnabled {
		g.Add(notificationsService.RunNewOrderNotifier, func(err error) {
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

	adminsRepo := admins.NewPg(db)
	adminsService := admins.New(
		adminsRepo,
		log.With(zap.String("service", "admins")),
		admins.DOSpacesConfig{
			Endpoint: config.DigitalOcean.Spaces.Endpoint,
			Bucket:   config.DigitalOcean.Spaces.Bucket,
			ID:       config.DigitalOcean.Spaces.Key,
			Secret:   config.DigitalOcean.Spaces.Secret,
		},
		&notificationsAdminAdapter{
			notifier: notificationsService,
		},
	)
	adminsHandler := admins.MakeHandler(adminsService, authMw)

	webhookService := webhooks.New(
		&adminWebhooksAdapter{admin: adminsService},
		&notificationsWebhooksAdapter{notifier: notificationsService},
		log.With(zap.String("service", "webhooks")),
		config.Bot.ID,
	)
	webhooksHandler := webhooks.MakeHandler(
		webhookService,
		log.With(zap.String("service", "webhooks_server")),
		config.TelegramWebhooks.SecretToken)

	r.Mount("/api/v1/public/products", productsHandler)
	r.Mount("/api/v1/public/auth", tgUsersHandler)
	r.Mount("/api/v1/public/orders", ordersHandler)
	r.Mount("/api/v1/private/marketplaces", adminsHandler)
	r.Mount("/api/v1/telegram/webhooks", webhooksHandler)

	g.Add(func() error {
		log.Info("starting HTTP server", zap.String("port", config.HTTP.Port))
		return httpServer.ListenAndServe()
	}, func(err error) {
		_ = httpServer.Shutdown(ctx)
	})

	if err := g.Run(); err != nil {
		log.Fatal("api exited with error:", logging.SilentError(err))
		return
	}
}
