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
)

// BookRepositoryInterface define contract for book related functions to repository
type BookRepositoryInterface interface {
	GetBooks(ctx context.Context) ([]*entity.Book, error)
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
