package entity

import (
	"time"
)

// Order struct holds entity of order
type Order struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	Fee        int       `json:"fee"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
