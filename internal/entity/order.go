package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/response"
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

// OrderPayload holds order payload representative
type OrderPayload struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

// Validate is func to validate order payload
func (o *OrderPayload) Validate() error {
	if o.Quantity <= 0 {
		return response.ErrInvalidQuantity
	}

	return nil
}
