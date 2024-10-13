package logger

import (
	"log"
)

func Error(message string, err error) {
	log.Printf("ERROR: %s: %v", message, err)
}

// Add other logging functions (Info, Warn, Debug) as needed
