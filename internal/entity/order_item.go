package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/response"
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
	Book           *Book     `json:"book"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// OrderItemPayload holds orderItem payload representative
type OrderItemPayload struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

// Validate is func to validate orderItem payload
func (o *OrderItemPayload) Validate() error {
	if o.Quantity <= 0 {
		return response.ErrInvalidQuantity
	}

	return nil
}
