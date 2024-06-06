package entity

import (
	"time"
)

// Order struct holds entity of order
type Order struct {
	ID         int64     `json:"id"`
	UserID     int32     `json:"user_id"`
	BookID     int32     `json:"book_id"`
	Quantity   int32     `json:"quantity"`
	Price      int32     `json:"price"`
	Fee        int32     `json:"fee"`
	TotalPrice int32     `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
