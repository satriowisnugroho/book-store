package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// UserRepositoryInterface define contract for user related functions to repository
type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
}

// UserRepository holds database connection
type UserRepository struct {
	db *sqlx.DB
}

var (
	// UserTableName hold table name for users
	UserTableName = "users"
	// UserColumns list all columns on users table
	UserColumns = []string{"id", "email", "fullname", "crypted_password", "created_at", "updated_at"}
	// UserAttributes hold string format of all users table columns
	UserAttributes = strings.Join(UserColumns, ", ")

	// UserCreationColumns list all columns used for create user
	UserCreationColumns = UserColumns[1:]
	// UserCreationAttributes hold string format of all creation user columns
	UserCreationAttributes = strings.Join(UserCreationColumns, ", ")
)

// NewUserRepository create initiate user repository with given database
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser insert user data into database
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	functionName := "UserRepository.CreateUser"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	// Handle case-insensitive
	user.Email = strings.ToLower(user.Email)

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, UserTableName, UserCreationAttributes, EnumeratedBindvars(UserCreationColumns))

	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Fullname,
		user.CryptedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
	if err != nil {
		if isUniqueConstraintViolation(err) {
			return response.ErrDuplicateEmail
		}

		return errors.Wrap(err, functionName)
	}

	return nil
}
