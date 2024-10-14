package logger

import (
	"log"
)

func Info(message string) {
	log.Printf("INFO: %s", message)
}

func Error(message string, err error) {
	log.Printf("ERROR: %s: %v", message, err)
}

// Add other logging functions (Warn, Debug) as needed
