package main

import (
	"log/slog"
	"os"
	"rest_server/config"
	"rest_server/handlers"
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
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", handlers.New(log, stg))

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
