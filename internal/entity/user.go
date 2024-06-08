package entity

import (
	"time"
)

// User struct holds entity of user
type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Fullname        string    `json:"fullname"`
	CryptedPassword int       `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
