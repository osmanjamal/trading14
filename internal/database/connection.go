package database

import (
	"database/sql"

	_ "github.com/lib/pq" // driver for PostgreSQL
)

func Connect(databaseURL string) (*DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return NewDB(db), nil
}
