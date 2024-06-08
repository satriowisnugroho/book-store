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
	dbentity "github.com/satriowisnugroho/book-store/internal/repository/postgres/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// UserRepositoryInterface define contract for user related functions to repository
type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
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

func (r *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.User, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.User, 0)

	for rows.Next() {
		tmpEntity := dbentity.User{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
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

// GetUserByEmail query to get user by email
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	functionName := "UserRepository.GetUserByEmail"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE email = $1 LIMIT 1", UserAttributes, UserTableName)
	rows, err := r.fetch(ctx, query, strings.ToLower(email))
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if len(rows) == 0 {
		return nil, response.ErrNotFound
	}

	return rows[0], nil
}
