package entity

import (
	"time"
)

// OrderItem struct holds entity of orderItem
type OrderItem struct {
	ID             int       `json:"id"`
	OrderID        int       `json:"order_id"`
	BookID         int       `json:"book_id"`
	Quantity       int       `json:"quantity"`
	Price          int       `json:"price"`
	Fee            int       `json:"fee"`
	TotalItemPrice int       `json:"total_item_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
