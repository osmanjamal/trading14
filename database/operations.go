package database

import (
	"time"
)

type Bot struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Strategy string `json:"strategy"`
	Status   string `json:"status"`
}

type Trade struct {
	ID        int       `json:"id"`
	BotID     int       `json:"bot_id"`
	Symbol    string    `json:"symbol"`
	Side      string    `json:"side"`
	Amount    float64   `json:"amount"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func (db *DB) GetBots() ([]Bot, error) {
	rows, err := db.Query("SELECT id, name, strategy, status FROM bots")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bots []Bot
	for rows.Next() {
		var b Bot
		if err := rows.Scan(&b.ID, &b.Name, &b.Strategy, &b.Status); err != nil {
			return nil, err
		}
		bots = append(bots, b)
	}
	return bots, nil
}

func (db *DB) GetBot(id int) (*Bot, error) {
	var b Bot
	err := db.QueryRow("SELECT id, name, strategy, status FROM bots WHERE id = $1", id).Scan(&b.ID, &b.Name, &b.Strategy, &b.Status)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (db *DB) CreateBot(bot *Bot) error {
	return db.QueryRow("INSERT INTO bots (name, strategy, status) VALUES ($1, $2, $3) RETURNING id", bot.Name, bot.Strategy, "stopped").Scan(&bot.ID)
}

func (db *DB) UpdateBot(bot *Bot) error {
	_, err := db.Exec("UPDATE bots SET name = $1, strategy = $2, status = $3 WHERE id = $4", bot.Name, bot.Strategy, bot.Status, bot.ID)
	return err
}

func (db *DB) DeleteBot(id int) error {
	_, err := db.Exec("DELETE FROM bots WHERE id = $1", id)
	return err
}

func (db *DB) GetTrades() ([]Trade, error) {
	rows, err := db.Query("SELECT id, bot_id, symbol, side, amount, price, timestamp FROM trades ORDER BY timestamp DESC LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []Trade
	for rows.Next() {
		var t Trade
		if err := rows.Scan(&t.ID, &t.BotID, &t.Symbol, &t.Side, &t.Amount, &t.Price, &t.Timestamp); err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, nil
}

func (db *DB) GetBotTrades(botID int) ([]Trade, error) {
	rows, err := db.Query("SELECT id, bot_id, symbol, side, amount, price, timestamp FROM trades WHERE bot_id = $1 ORDER BY timestamp DESC LIMIT 100", botID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []Trade
	for rows.Next() {
		var t Trade
		if err := rows.Scan(&t.ID, &t.BotID, &t.Symbol, &t.Side, &t.Amount, &t.Price, &t.Timestamp); err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, nil
}

func (db *DB) StartBot(id int) error {
	_, err := db.Exec("UPDATE bots SET status = 'running' WHERE id = $1", id)
	return err
}

func (db *DB) StopBot(id int) error {
	_, err := db.Exec("UPDATE bots SET status = 'stopped' WHERE id = $1", id)
	return err
}
