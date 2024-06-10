package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/entity"
)

// OrderItem struct holds orderItem database representative
type OrderItem struct {
	ID             int       `db:"id"`
	OrderID        int       `db:"order_id"`
	BookID         int       `db:"book_id"`
	Quantity       int       `db:"quantity"`
	Price          int       `db:"price"`
	TotalItemPrice int       `db:"total_item_price"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// ToEntity to convert orderItem from database to entity contract
func (e *OrderItem) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ID:             e.ID,
		OrderID:        e.OrderID,
		BookID:         e.BookID,
		Quantity:       e.Quantity,
		Price:          e.Price,
		TotalItemPrice: e.TotalItemPrice,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}
