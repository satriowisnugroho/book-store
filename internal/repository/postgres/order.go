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

// OrderRepositoryInterface define contract for order related functions to repository
type OrderRepositoryInterface interface {
	GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)
}

// OrderRepository holds database connection
type OrderRepository struct {
	db *sqlx.DB
}

var (
	// OrderTableName hold table name for orders
	OrderTableName = "orders"
	// OrderColumns list all columns on orders table
	OrderColumns = []string{"id", "user_id", "book_id", "quantity", "price", "fee", "total_price", "created_at", "updated_at"}
	// OrderAttributes hold string format of all orders table columns
	OrderAttributes = strings.Join(OrderColumns, ", ")
)

// NewOrderRepository create initiate order repository with given database
func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Order, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Order, 0)

	for rows.Next() {
		tmpEntity := dbentity.Order{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
}

// GetOrders query to get list of orders by user ID
func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	functionName := "OrderRepository.GetOrdersByUserID"
	if err := helper.CheckDeadline(ctx); err != nil {
		return []*entity.Order{}, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = %d", OrderAttributes, OrderTableName, userID)

	rows, err := r.fetch(ctx, query)
	if err != nil {
		return rows, errors.Wrap(err, functionName)
	}

	return rows, nil
}
