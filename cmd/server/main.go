package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/osmanjamal/trading14/internal/api"
	"github.com/osmanjamal/trading14/internal/config"
	"github.com/osmanjamal/trading14/internal/database"
	"github.com/osmanjamal/trading14/pkg/logger"
)

func main() {
	// تحميل التكوين
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// إنشاء المسجل
	logger := logger.New(cfg.LogLevel)
	defer logger.Sync()

	// الاتصال بقاعدة البيانات
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}
	defer db.Close()

	// إنشاء الراوتر
	r := mux.NewRouter()

	// إنشاء معالجات API
	handlers := api.NewHandlers(db, logger)

	// إعداد مسارات API
	r.HandleFunc("/api/bots", handlers.GetBots).Methods("GET")
	r.HandleFunc("/api/bots/{id}", handlers.GetBot).Methods("GET")
	r.HandleFunc("/api/bots", handlers.CreateBot).Methods("POST")
	r.HandleFunc("/api/bots/{id}", handlers.UpdateBot).Methods("PUT")
	r.HandleFunc("/api/bots/{id}", handlers.DeleteBot).Methods("DELETE")
	r.HandleFunc("/api/bots/{id}/start", handlers.StartBot).Methods("POST")
	r.HandleFunc("/api/bots/{id}/stop", handlers.StopBot).Methods("POST")
	r.HandleFunc("/api/trades", handlers.GetTrades).Methods("GET")
	r.HandleFunc("/api/bots/{id}/trades", handlers.GetBotTrades).Methods("GET")

	// إنشاء خادم HTTP
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// تشغيل الخادم في goroutine منفصلة
	go func() {
		logger.Info("Starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed", "error", err)
		}
	}()

	// إعداد قناة لإشارات الإيقاف
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// إغلاق الخادم بشكل آمن
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exiting")
}
