package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/entity"
)

// Book struct holds book database representative
type Book struct {
	ID        int       `db:"id"`
	Isbn      string    `db:"isbn"`
	Title     string    `db:"title"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ToEntity to convert book from database to entity contract
func (e *Book) ToEntity() *entity.Book {
	return &entity.Book{
		ID:        e.ID,
		Isbn:      e.Isbn,
		Title:     e.Title,
		Price:     e.Price,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
