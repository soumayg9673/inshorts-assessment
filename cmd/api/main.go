package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/soumayg9673/inshorts-assessment/internal/env"
	"github.com/soumayg9673/inshorts-assessment/internal/middleware"
	"go.uber.org/zap"
)

func init() {
	// load .env File
	godotenv.Load(".env")
}

func main() {

	// Initialize configuration
	cfg := config{
		addr:    env.GetString("ADDR", ":8080"),
		env:     env.GetString("ENV", "local"),
		version: env.GetString("VERSION", "1.0.0"),
		db: dbConfig{
			addr:         env.GetString("DATABASE_URL", ""),
			maxOpenConns: env.GetInt("DATABASE_MAX_OPEN_CONNS", 10),
			maxIdleConns: env.GetInt("DATABASE_MAX_IDLE_CONNS", 10),
			maxIdleTime:  env.GetString("DATABASE_MAX_OPEN_TIME", "5m"),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction())
	if env.GetString("ENV", "local") == "local" {
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()

	middleware := middleware.NewMiddleware(cfg.env, logger)

	app := &application{
		config:     cfg,
		logger:     logger,
		middleware: middleware,
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatalln(err.Error())
	}
}
