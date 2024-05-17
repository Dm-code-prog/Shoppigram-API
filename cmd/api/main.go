package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shoppigram-com/marketplace-api/internal/admins"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"github.com/shoppigram-com/marketplace-api/internal/notifications"
	"github.com/shoppigram-com/marketplace-api/internal/orders"
	"go.uber.org/zap/zapcore"

	"github.com/Netflix/go-env"
	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oklog/run"
	"github.com/shoppigram-com/marketplace-api/internal/cors"
	"github.com/shoppigram-com/marketplace-api/internal/products"
	productsgenerated "github.com/shoppigram-com/marketplace-api/internal/products/generated"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
)

func main() {
	var (
		logLevel  zapcore.Level
		zapConfig zap.Config
	)

	envLogLevel := os.Getenv("LOG_LEVEL")

	switch envLogLevel {
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

	if envLogLevel == "DEBUG" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Level.SetLevel(logLevel)

	log, _ := zapConfig.Build(zap.AddStacktrace(zapcore.PanicLevel))
	defer log.Sync()

	ctx := context.Background()

	var config Environment
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal("failed to load environment variables", logging.SilentError(err))
		return
	}

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
		cors.Middleware,
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

	productsRepo := products.NewPg(productsgenerated.New(db))
	productsService := products.New(productsRepo, log.With(zap.String("service", "products")), productsCache)
	productsHandler := products.MakeHandler(productsService)

	tgUsersRepo := telegramusers.NewPg(db, config.Encryption.Key)
	tgUsersService := telegramusers.New(tgUsersRepo, log.With(zap.String("service", "users")))
	tgUsersHandler := telegramusers.MakeHandler(tgUsersService, authMw)

	ordersRepo := orders.NewPg(db)
	ordersService := orders.New(ordersRepo, log.With(zap.String("service", "orders")))
	ordersHandler := orders.MakeHandler(ordersService, authMw)

	adminsRepo := admins.NewPg(db)
	adminsService := admins.New(adminsRepo, log.With(zap.String("service", "admins")))
	adminsHandler := admins.MakeHandler(adminsService, authMw)

	notificationsRepo := notifications.NewPg(
		db,
		config.NewOrderNotifications.BatchSize,
		config.NewMarketplaceNotifications.BatchSize,
	)
	notificationsService := notifications.New(
		notificationsRepo,
		log.With(zap.String("service", "notifications")),
		time.Duration(config.NewOrderNotifications.Timeout)*time.Second,
		time.Duration(config.NewMarketplaceNotifications.Timeout)*time.Second,
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

	r.Mount("/api/v1/public/products", productsHandler)
	r.Mount("/api/v1/public/auth", tgUsersHandler)
	r.Mount("/api/v1/public/orders", ordersHandler)
	r.Mount("/api/v1/private/marketplaces", adminsHandler)

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
