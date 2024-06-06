package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/entity"
)

// Order struct holds order database representative
type Order struct {
	ID         int64     `db:"id"`
	UserID     int32     `db:"user_id"`
	BookID     int32     `db:"book_id"`
	Quantity   int32     `db:"quantity"`
	Price      int32     `db:"price"`
	Fee        int32     `db:"fee"`
	TotalPrice int32     `db:"total_price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// ToEntity to convert order from database to entity contract
func (e *Order) ToEntity() *entity.Order {
	return &entity.Order{
		ID:         e.ID,
		UserID:     e.UserID,
		BookID:     e.BookID,
		Quantity:   e.Quantity,
		Price:      e.Price,
		Fee:        e.Fee,
		TotalPrice: e.TotalPrice,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}
