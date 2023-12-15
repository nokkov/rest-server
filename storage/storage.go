package database

import "database/sql"

type Entity struct {
	id        int
	url       string
	short_url string
}

type Storage struct {
	db *sql.DB
}
