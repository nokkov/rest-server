package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	EnvCfg    EnvConfig
	DbCfg     DatabaseConfig
	ServerCfg HttpServerConfig
}

type EnvConfig struct {
	Env string `yaml: "env" env-default: "local" env-required: "true"`
}

type DatabaseConfig struct {
	Type     string `yaml: "type" env-required: "true"`
	Host     string `yaml: "host" env-required: "true"`
	Port     string `yaml: "port" env-required: "true"`
	User     string `yaml: "user" env-required: "true"`
	Password string `yaml: "password" env-required: "true"`
	DbName   string `yaml: "dbname" env-required: "true"`
}

type HttpServerConfig struct {
	Address     string        `yaml: "address" env-required: "true"`
	IdleTimeout time.Duration `yaml: "idle_timeout" env-required: "true"`
}

func MustLoad() AppConfig {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var envCfg EnvConfig

	var dbCfg DatabaseConfig

	var serverCfg HttpServerConfig

	if err := cleanenv.ReadConfig(configPath, &envCfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	if err := cleanenv.ReadConfig(configPath, &dbCfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	if err := cleanenv.ReadConfig(configPath, &serverCfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return AppConfig{envCfg, dbCfg, serverCfg}
}
