package api

import (
	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/pkg/logger"
)

func SetupRoutes(r *mux.Router, db *database.DB, logger *logger.Logger) {
	h := NewHandlers(db, logger)

	r.HandleFunc("/api/bots", h.GetBots).Methods("GET")
	r.HandleFunc("/api/bots", h.CreateBot).Methods("POST")
	r.HandleFunc("/api/trades", h.GetTrades).Methods("GET")
}
