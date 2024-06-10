package entity

import (
	"time"
)

// Order struct holds entity of order
type Order struct {
	ID         int          `json:"id"`
	UserID     int          `json:"user_id"`
	Fee        int          `json:"fee"`
	TotalPrice int          `json:"total_price"`
	OrderItems []*OrderItem `json:"order_items"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

// OrderPayload holds order payload representative
type OrderPayload struct {
	OrderItems []OrderItemPayload `json:"order_items"`
}

// OrderResponse holds order response
type OrderResponse struct {
	*Order
	Book *Book
}
