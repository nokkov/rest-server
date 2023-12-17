package main

import (
	"fmt"
	"log/slog"
	"os"
	"rest_server/config"
	database "rest_server/storage"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()
	log := setupLogger(config.EnvCfg.Env)

	log.Info("starting url-shortener...")

	stg, err := database.New(&config.DbCfg)

	if err != nil {
		fmt.Print(err)
	}

	log.Info("database create...")

	link, err := stg.GetUrl("short.devops/google") //is it actually short? xD

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(link)
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
