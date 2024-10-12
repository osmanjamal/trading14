package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/yourusername/trading-bot/internal/api"
    "github.com/yourusername/trading-bot/internal/bot"
    "github.com/yourusername/trading-bot/internal/database"
    "github.com/yourusername/trading-bot/internal/exchange"
    "github.com/yourusername/trading-bot/pkg/logger"
)

func TestHandleWebhook(t *testing.T) {
    // Mock dependencies
    mockDB := &database.MockDB{}
    mockExchange := &exchange.MockExchange{}
    mockLogger := logger.New("debug")

    // Create handlers
    handlers := api.NewHandlers(mockDB, mockExchange, mockLogger)

    // Create a test signal
    signal := bot.Signal{
        Symbol: "BTCUSDT",
        Action: "buy",
        Price:  50000,
    }

    // Create a request
    body, _ := json.Marshal(signal)
    req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler
    handlers.HandleWebhook(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check that the mock exchange was called
    if !mockExchange.PlaceOrderCalled {
        t.Error("PlaceOrder was not called on the exchange")
    }

    // Check that the mock DB was called
    if !mockDB.LogTradeCalled {
        t.Error("LogTrade was not called on the database")
    }
}