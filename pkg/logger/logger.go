package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(level string) *Logger {
	var zapLogger *zap.Logger
	var err error

	switch level {
	case "debug":
		zapLogger, err = zap.NewDevelopment()
	default:
		zapLogger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	return &Logger{zapLogger.Sugar()}
}

// تأكد من إضافة هذه الأساليب لتغطية جميع مستويات التسجيل
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Infow(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
}
