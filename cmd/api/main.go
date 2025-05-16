package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/soumayg9673/inshorts-assessment/internal/database"
	"github.com/soumayg9673/inshorts-assessment/internal/env"
	"github.com/soumayg9673/inshorts-assessment/internal/llm"
	"github.com/soumayg9673/inshorts-assessment/internal/middleware"
	"github.com/soumayg9673/inshorts-assessment/internal/repository"
	"github.com/soumayg9673/inshorts-assessment/internal/service"
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

	// Database
	db, err := database.NewDatabaseConn(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
		cfg.env,
	)
	if err != nil {
		logger.Panic(err.Error(),
			zap.String("env", cfg.env),
		)
	}
	defer db.Close()
	logger.Info("connected to the database with connection pooling",
		zap.String("env", cfg.env),
	)

	llm := llm.NewLlmStore(logger)

	dtb := database.NewDbStore(db, logger, cfg.env)
	rpo := repository.NewRpoStore(dtb, logger, cfg.env)
	svc := service.NewServiceStore(rpo, logger, llm, cfg.env)

	middleware := middleware.NewMiddleware(cfg.env, logger)

	app := &application{
		config:     cfg,
		service:    svc,
		logger:     logger,
		middleware: middleware,
	}

	if len(os.Args) > 1 && os.Args[1] == "insert" {
		insertInitialData(db)
		logger.Info("data insert successful",
			zap.String("env", cfg.env),
		)
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		logger.Fatal(err.Error(),
			zap.String("env", cfg.env),
		)
	}
}
