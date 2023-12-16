package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func (stg *Storage) SaveUrl(urlToSave, short_url string) error {
	stmt, err := stg.db.Prepare("INSERT INTO urls(url, short_url) VALUES ($1, $2)")

	if err != nil {
		return fmt.Errorf("SaveUrl() prepare: %s", err)
	}

	_, err = stmt.Exec(urlToSave, short_url)
	if err != nil {
		return fmt.Errorf("SaveUrl() exec prepare: %s", err)
	}

	return nil
}
