package auth

import "golang.org/x/crypto/bcrypt"

// PasswordHasher defines an interface for password hashing
type PasswordHasher interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

// BcryptPasswordHasher is the default implementation using bcrypt
type BcryptPasswordHasher struct{}

// GenerateFromPassword generates a bcrypt hash of the password
func (h *BcryptPasswordHasher) GenerateFromPassword(password string) (string, error) {
	byteCryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byteCryptedPassword), err
}

// GenerateFromPassword compares a bcrypt hashed password with its possible plaintext equivalent
func (h *BcryptPasswordHasher) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
