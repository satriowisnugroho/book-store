package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	dbentity "github.com/satriowisnugroho/book-store/internal/repository/postgres/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// BookRepositoryInterface define contract for book related functions to repository
type BookRepositoryInterface interface {
	GetBooks(ctx context.Context) ([]*entity.Book, error)
	GetBooksCount(ctx context.Context) (int, error)
	GetBookByID(ctx context.Context, bookID int) (*entity.Book, error)
}

// BookRepository holds database connection
type BookRepository struct {
	db *sqlx.DB
}

var (
	// BookTableName hold table name for books
	BookTableName = "books"
	// BookColumns list all columns on books table
	BookColumns = []string{"id", "isbn", "title", "price", "created_at", "updated_at"}
	// BookAttributes hold string format of all books table columns
	BookAttributes = strings.Join(BookColumns, ", ")
)

// NewBookRepository create initiate book repository with given database
func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Book, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Book, 0)

	for rows.Next() {
		tmpEntity := dbentity.Book{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
}

// GetBooks query to get list of books
func (r *BookRepository) GetBooks(ctx context.Context) ([]*entity.Book, error) {
	functionName := "BookRepository.GetBooks"

	if err := helper.CheckDeadline(ctx); err != nil {
		return []*entity.Book{}, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s", BookAttributes, BookTableName)
	rows, err := r.fetch(ctx, query)
	if err != nil {
		return rows, errors.Wrap(err, functionName)
	}

	return rows, nil
}

// GetBooksCount query to get the count of books
func (r *BookRepository) GetBooksCount(ctx context.Context) (int, error) {
	functionName := "BookRepository.GetBooksCount"
	if err := helper.CheckDeadline(ctx); err != nil {
		return 0, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", BookTableName)

	count := 0
	rows := r.db.QueryRowxContext(ctx, query)
	if err := rows.Scan(&count); err != nil {
		return count, errors.Wrap(err, functionName)
	}

	return count, nil
}

// GetBookByID query to get book by ID
func (r *BookRepository) GetBookByID(ctx context.Context, bookID int) (*entity.Book, error) {
	functionName := "BookRepository.GetBookByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1 LIMIT 1", BookAttributes, BookTableName)
	rows, err := r.fetch(ctx, query, bookID)
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if len(rows) == 0 {
		return nil, response.ErrNotFound
	}

	return rows[0], nil
}
