package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/pkg/logger"
)

type Handlers struct {
	db     *database.DB
	logger *logger.Logger
}

func NewHandlers(db *database.DB, logger *logger.Logger) *Handlers {
	return &Handlers{db: db, logger: logger}
}

func (h *Handlers) GetBots(w http.ResponseWriter, r *http.Request) {
	bots, err := h.db.GetBots()
	if err != nil {
		h.logger.Error("Failed to get bots", "error", err)
		http.Error(w, "Failed to get bots", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bots)
}

func (h *Handlers) GetBot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	bot, err := h.db.GetBot(id)
	if err != nil {
		h.logger.Error("Failed to get bot", "error", err, "id", id)
		http.Error(w, "Bot not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bot)
}

func (h *Handlers) CreateBot(w http.ResponseWriter, r *http.Request) {
	var bot database.Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.CreateBot(&bot); err != nil {
		h.logger.Error("Failed to create bot", "error", err)
		http.Error(w, "Failed to create bot", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bot)
}

func (h *Handlers) UpdateBot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	var bot database.Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bot.ID = id

	if err := h.db.UpdateBot(&bot); err != nil {
		h.logger.Error("Failed to update bot", "error", err, "id", id)
		http.Error(w, "Failed to update bot", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bot)
}

func (h *Handlers) DeleteBot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteBot(id); err != nil {
		h.logger.Error("Failed to delete bot", "error", err, "id", id)
		http.Error(w, "Failed to delete bot", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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

func (h *Handlers) GetBotTrades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	botID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	trades, err := h.db.GetBotTrades(botID)
	if err != nil {
		h.logger.Error("Failed to get bot trades", "error", err, "botID", botID)
		http.Error(w, "Failed to get bot trades", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(trades)
}

func (h *Handlers) StartBot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	if err := h.db.StartBot(id); err != nil {
		h.logger.Error("Failed to start bot", "error", err, "id", id)
		http.Error(w, "Failed to start bot", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Bot started"})
}

func (h *Handlers) StopBot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bot ID", http.StatusBadRequest)
		return
	}

	if err := h.db.StopBot(id); err != nil {
		h.logger.Error("Failed to stop bot", "error", err, "id", id)
		http.Error(w, "Failed to stop bot", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Bot stopped"})
}
