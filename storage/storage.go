package database

import (
	"database/sql"
	"fmt"
	"log"
	"rest_server/config"
)

type Entity struct {
	id        int
	url       string
	short_url string
}

type Storage struct {
	db *sql.DB
}

// default database is postgre => planning to make it possible to swotch
// databases
func New(cfg *config.DatabaseConfig, storage *Storage) (*Storage, error) {

	var driver string
	var connInfo string

	switch cfg.Type {
	case "postgresql":
		driver = "postgresql"
		connInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	case "mysql":
		driver = "mysql"
		connInfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	default:
		log.Fatalf("unknown database type: %s", cfg.Type)
	}

	db, err := sql.Open(driver, connInfo)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", connInfo, err)
	}

	return &Storage{db: db}, nil
}
