package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Trade struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}
