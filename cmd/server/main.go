package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/osmanjamal/trading12/internal/api"
	"github.com/osmanjamal/trading12/internal/config"
	"github.com/osmanjamal/trading12/internal/database"
	"github.com/osmanjamal/trading12/internal/exchange"
	"github.com/osmanjamal/trading12/pkg/logger"
)

func main() {
	// تحميل التكوين
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("خطأ في تحميل التكوين: %v", err)
	}

	// إعداد المسجل
	logger := logger.New(cfg.LogLevel)
	defer logger.Sync()

	// الاتصال بقاعدة البيانات
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("فشل الاتصال بقاعدة البيانات", "خطأ", err)
	}
	defer db.Close()

	// اختبار الاتصال بقاعدة البيانات
	if err = db.Ping(); err != nil {
		logger.Fatal("فشل اختبار الاتصال بقاعدة البيانات", "خطأ", err)
	}
	logger.Info("تم الاتصال بقاعدة البيانات بنجاح")

	// إعداد اتصال منصة التداول
	exchange := exchange.NewBinance(cfg.ExchangeAPIKey, cfg.ExchangeSecretKey)

	// إعداد الموجه والمعالجات
	router := api.SetupRoutes(db, exchange, logger)

	// إنشاء خادم HTTP
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// تشغيل الخادم في goroutine منفصلة
	go func() {
		logger.Info("بدء تشغيل الخادم", "منفذ", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("فشل الخادم", "خطأ", err)
		}
	}()

	// إعداد قناة لإشارات إيقاف التشغيل
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("جاري إيقاف تشغيل الخادم...")

	// إغلاق الخادم بشكل آمن
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("تم إجبار الخادم على الإغلاق", "خطأ", err)
	}

	logger.Info("تم إيقاف تشغيل الخادم بنجاح")
}
