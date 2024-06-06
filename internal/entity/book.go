package entity

import (
	"time"
)

// Book struct holds entity of book
type Book struct {
	ID        int64     `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Price     int32     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
