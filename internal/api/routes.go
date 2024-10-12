package api

import (
	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading12/internal/database"
	"github.com/osmanjamal/trading12/internal/exchange"
	"github.com/osmanjamal/trading12/pkg/logger"
)

func SetupRoutes(db *database.DB, exchange exchange.Exchange, logger *logger.Logger) *mux.Router {
	r := mux.NewRouter()

	h := NewHandlers(db, exchange, logger)

	r.HandleFunc("/webhook", h.HandleWebhook).Methods("POST")
	r.HandleFunc("/api/trades", h.GetTrades).Methods("GET")
	r.HandleFunc("/api/balance", h.GetBalance).Methods("GET")

	return r
}
