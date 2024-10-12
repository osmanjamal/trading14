package database

import (
	"time"
)

type Trade struct {
	ID        int64     `json:"id"`
	Symbol    string    `json:"symbol"`
	Action    string    `json:"action"`
	Price     float64   `json:"price"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
