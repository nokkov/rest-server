package config

type AppConfig struct {
	Env string `yaml: "env" env-default: "local" env-required: "true"`
}

type DatabaseConfig struct {
	Host     string `yaml: "host" env-required: "true"`
	Port     string `yaml: "port" env-required: "true"`
	User     string `yaml: "user" env-required: "true"`
	Password string `yaml: "password" env-required: "true"`
	DbName   string `yaml: "dbname" env-required: "true"`
}

type HttpServerConfig struct {
	Address     string `yaml: "address" env-required: "true"`
	IdleTimeout int    `yaml: "idle_timeout" env-required: "true"`
}
