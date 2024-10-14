package database

import (
	"database/sql"
	"time"
)

type DB struct {
	*sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db}
}

func (db *DB) LogTrade(symbol string, action string, price float64, amount float64) error {
	_, err := db.Exec(`
        INSERT INTO trades (symbol, action, price, amount, timestamp)
        VALUES ($1, $2, $3, $4, $5)
    `, symbol, action, price, amount, time.Now())
	return err
}
