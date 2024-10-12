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