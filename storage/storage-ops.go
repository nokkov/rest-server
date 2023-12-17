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

	defer stmt.Close()

	_, err = stmt.Exec(urlToSave, short_url)
	if err != nil {
		return fmt.Errorf("SaveUrl() exec prepare: %s", err)
	}

	return nil
}

func (stg *Storage) GetUrl(searchedUrl string) (string, error) {
	stmt, err := stg.db.Prepare("SELECT * FROM urls WHERE short_url = $1")

	if err != nil {
		return "", fmt.Errorf("GetUrl() prepare: %s", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(searchedUrl)

	if row.Err() != nil {
		return "", fmt.Errorf("GetUrl() query row: %s", err)
	}

	var result *Entity

	row.Scan(result)

	if result.url != "" {
		return result.url, nil
	} else {
		return "", nil
	}
}
