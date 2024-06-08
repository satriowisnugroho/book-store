package entity

import (
	"regexp"
	"time"

	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/response"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// User struct holds entity of user
type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Fullname        string    `json:"fullname"`
	CryptedPassword string    `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// UserPayload holds order payload representative
type UserPayload struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

// Validate is func to validate order payload
func (u *UserPayload) Validate() error {
	if !emailRegex.MatchString(u.Email) {
		return response.ErrInvalidEmail
	}

	if len(u.Fullname) == 0 {
		return response.ErrInvalidFullname
	}

	if len(u.Password) < config.MinPasswordLen {
		return response.ErrInvalidPasswordLength
	}

	return nil
}
