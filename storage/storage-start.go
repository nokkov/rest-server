package storage

import (
	"database/sql"
	"fmt"
	"log"
	"rest_server/config"

	_ "github.com/lib/pq"
)

type Entity struct {
	id        int
	url       string
	short_url string
}

func New(cfg *config.DatabaseConfig) (*Storage, error) {
	var driver string
	var connInfo string

	switch cfg.Type {
	case "postgresql":
		driver = "postgres"
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

	err = MustLoad(db)

	if err != nil {
		return nil, fmt.Errorf("ping error: %s", err)
	}

	return &Storage{db: db}, nil
}

func MustLoad(db *sql.DB) error {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			url TEXT NOT NULL UNIQUE,
			short_url TEXT NOT NULL UNIQUE
		);
	`)
	defer stmt.Close()

	if err != nil {
		return fmt.Errorf("MustLoad() prepare: %s", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("MustLoad() exec prepare: %s", err)
	}

	return db.Ping()
}
