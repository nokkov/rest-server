package main

import (
	"log/slog"
	"net/http"
	"os"
	"rest_server/config"
	handlers "rest_server/controllers"
	"rest_server/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()
	log := setupLogger(config.EnvCfg.Env)

	log.Info("starting url-shortener...")

	stg, _ := storage.New(&config.DbCfg)

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // useful for tracing
	router.Use(middleware.Recoverer) // if a panic on the server => generate http-status 500
	router.Use(middleware.URLFormat) // parse url extension and store it on the context
	//TODO: custom error router
	router.Post("/save-url", handlers.New(log, stg))
	router.Get("/get-base-url/{short_url}", handlers.Get(log, stg))

	log.Info("starting server...", slog.String("address", config.ServerCfg.Address))

	srv := &http.Server{
		Addr:         config.ServerCfg.Address,
		Handler:      router,
		ReadTimeout:  config.ServerCfg.Timeout,
		WriteTimeout: config.ServerCfg.Timeout,
		IdleTimeout:  config.ServerCfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	//srv.ListenAndServe() is a "blocking" function, if the code went beyond it, an error occurred
	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
