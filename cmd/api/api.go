package main

import (
	"net/http"
	"time"

	"github.com/soumayg9673/inshorts-assessment/internal/middleware"
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

	return srv.ListenAndServe()
}
