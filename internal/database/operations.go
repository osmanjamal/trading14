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

func (db *DB) GetTrades() ([]Trade, error) {
	rows, err := db.Query("SELECT id, symbol, action, price, amount, timestamp FROM trades ORDER BY timestamp DESC LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []Trade
	for rows.Next() {
		var t Trade
		err := rows.Scan(&t.ID, &t.Symbol, &t.Action, &t.Price, &t.Amount, &t.Timestamp)
		if err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, nil
}
