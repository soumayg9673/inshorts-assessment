package main

import (
	"net/http"
	"time"

	"github.com/soumayg9673/inshorts-assessment/internal/middleware"
	"github.com/soumayg9673/inshorts-assessment/internal/service"
	"go.uber.org/zap"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr    string
	db      dbConfig
	env     string
	version string
}

type application struct {
	config     config
	service    service.Service
	logger     *zap.Logger
	middleware middleware.Middleware
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	api := http.NewServeMux()
	// TODO: Register API routes
	mux.Handle("/api/", http.StripPrefix("/api", api))

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	stack := middleware.CreateStack(
		app.middleware.Middleware.LoggingMiddleware,
	)

	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      stack(mux),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 30,
	}

	app.logger.Info("server is up and running",
		zap.String("addr", app.config.addr),
		zap.String("env", app.config.env),
	)

	return srv.ListenAndServe()
}
