package tests

import (
    "testing"

    "github.com/yourusername/trading-bot/internal/bot"
    "github.com/yourusername/trading-bot/internal/database"
    "github.com/yourusername/trading-bot/internal/exchange"
)

func TestProcessSignal(t *testing.T) {
    // Mock dependencies
    mockDB := &database.MockDB{}
    mockExchange := &exchange.MockExchange{}

    // Create a test signal
    signal := bot.Signal{
        Symbol: "BTCUSDT",
        Action: "buy",
        Price:  50000,
    }

    // Process the signal
    err := bot.ProcessSignal(signal, mockExchange, mockDB)
    if err != nil {
        t.Fatalf("ProcessSignal returned an error: %v", err)
    }

    // Check that the mock exchange was called
    if !mockExchange.PlaceOrderCalled {
        t.Error("PlaceOrder was not called on the exchange")
    }

    // Check that the mock DB was called
    if !mockDB.LogTradeCalled {
        t.Error("LogTrade was not called on the database")
    }

    // Check the values passed to the exchange
    if mockExchange.LastSymbol != "BTCUSDT" || mockExchange.LastAction != "buy" || mockExchange.LastPrice != 50000 {
        t.Error("Incorrect values passed to the exchange")
    }

    // Check the values passed to the database
    if mockDB.LastSymbol != "BTCUSDT" || mockDB.LastAction != "buy" || mockDB.LastPrice != 50000 {
        t.Error("Incorrect values passed to the database")
    }
}