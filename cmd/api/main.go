package main

import (
	"context"
	"encoding/json"
	"github.com/Netflix/go-env"
	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/oklog/run"
	"github.com/shoppigram-com/marketplace-api/internal/products"
	"github.com/shoppigram-com/marketplace-api/internal/products/generated"
	"github.com/streadway/handy/cors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	ctx := context.Background()

	var config Environment
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal("failed to load environment variables", zap.Error(err))
	}

	db, err := pgx.Connect(ctx, config.Postgres.DSN)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close(ctx)
	log.Debug("connected to database")

	var g run.Group
	var r = chi.NewRouter()
	var httpServer = http.Server{
		Addr:    ":" + config.HTTP.Port,
		Handler: r,
	}

	r.Use(
		middleware.Timeout(10*time.Second),
		middleware.Recoverer,
		middleware.Compress(5, "application/json"),
		cors.Middleware(cors.Config{
			AllowOrigin: func(r *http.Request) string {
				return "*.shoppigram.com,*.shoppigram.dev,*.shoppigram.ru,localhost:3000"
			},
		}),
		middleware.Throttle(500),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "the path you requested does not exist"})
	})

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,         // number of keys to track frequency of (10M).
		MaxCost:     200_000_000, // maximum cost of cache (200 MB).
		BufferItems: 64,          // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("failed to create cache", zap.Error(err))
	}

	productsRepo := products.NewPg(generated.New(db))
	productsService := products.New(productsRepo, log.With(zap.String("service", "products")), cache)
	productsHandler := products.MakeHandler(productsService)

	r.Mount("/api/v1/public/products", productsHandler)

	g.Add(func() error {
		log.Info("starting HTTP server", zap.String("port", config.HTTP.Port))
		return httpServer.ListenAndServe()
	}, func(err error) {
		_ = httpServer.Shutdown(ctx)
	})

	if err := g.Run(); err != nil {
		log.Fatal("api exited with error:", zap.Error(err))
	}
}
