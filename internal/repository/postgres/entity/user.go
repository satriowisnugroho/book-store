package entity

import (
	"time"

	"github.com/satriowisnugroho/book-store/internal/entity"
)

// User struct holds user database representative
type User struct {
	ID              int       `db:"id"`
	Email           string    `db:"email"`
	Fullname        string    `db:"fullname"`
	CryptedPassword int       `db:"crypted_password"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

// ToEntity to convert user from database to entity contract
func (e *User) ToEntity() *entity.User {
	return &entity.User{
		ID:              e.ID,
		Email:           e.Email,
		Fullname:        e.Fullname,
		CryptedPassword: e.CryptedPassword,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}
