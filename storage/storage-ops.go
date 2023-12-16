package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func (stg *Storage) SaveUrl(urlToSave string, short_url string) error {
	stmt, err := stg.db.Prepare("INSERT INTO url(url, short_url) VALUES (?, ?)")

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	_, err = stmt.Exec(urlToSave, short_url)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
