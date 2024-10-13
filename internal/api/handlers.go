package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/internal/models"
	"github.com/osmanjamal/trading14/pkg/logger"
)

func InitDB() {
	err := database.Connect()
	if err != nil {
		logger.Error("Failed to connect to database", err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := database.GetUserByID(userID)
	if err != nil {
		logger.Error("Error fetching user data", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error("Invalid request payload", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = database.CreateUser(&user)
	if err != nil {
		logger.Error("Error creating user", err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error("Invalid request payload", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = database.UpdateUser(userID, &user)
	if err != nil {
		logger.Error("Error updating user", err)
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
