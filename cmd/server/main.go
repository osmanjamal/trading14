package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading14/internal/api"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/pkg/logger"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := mux.NewRouter()
	api.SetupRoutes(r)

	logger.Info("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
