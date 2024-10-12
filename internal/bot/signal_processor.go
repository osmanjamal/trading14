package bot

import (
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/internal/exchange"
)

type Signal struct {
	Symbol string  `json:"symbol"`
	Action string  `json:"action"`
	Price  float64 `json:"price"`
}

func ProcessSignal(signal Signal, exchange exchange.Exchange, db *database.DB) error {
	// Implement signal processing logic here
	// This is a simplified example
	var err error
	switch signal.Action {
	case "buy":
		err = exchange.PlaceOrder(signal.Symbol, "buy", signal.Price, 1.0)
	case "sell":
		err = exchange.PlaceOrder(signal.Symbol, "sell", signal.Price, 1.0)
	}

	if err != nil {
		return err
	}

	// Log the trade in the database
	return db.LogTrade(signal.Symbol, signal.Action, signal.Price, 1.0)
}
