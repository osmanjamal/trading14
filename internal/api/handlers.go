package api

import (
	"encoding/json"
	"net/http"

	"github.com/osmanjamal/trading14/internal/bot"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/internal/exchange"
	"github.com/osmanjamal/trading14/pkg/logger"
)

type Handlers struct {
	db       *database.DB
	exchange exchange.Exchange
	logger   *logger.Logger
}

func NewHandlers(db *database.DB, exchange exchange.Exchange, logger *logger.Logger) *Handlers {
	return &Handlers{db: db, exchange: exchange, logger: logger}
}

func (h *Handlers) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var signal bot.Signal
	if err := json.NewDecoder(r.Body).Decode(&signal); err != nil {
		h.logger.Error("Failed to decode webhook", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := bot.ProcessSignal(signal, h.exchange, h.db); err != nil {
		h.logger.Error("Failed to process signal", "error", err)
		http.Error(w, "Failed to process signal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) GetTrades(w http.ResponseWriter, r *http.Request) {
	trades, err := h.db.GetTrades()
	if err != nil {
		h.logger.Error("Failed to get trades", "error", err)
		http.Error(w, "Failed to get trades", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trades)
}

func (h *Handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	balance, err := h.exchange.GetBalance()
	if err != nil {
		h.logger.Error("Failed to get balance", "error", err)
		http.Error(w, "Failed to get balance", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(balance)
}
