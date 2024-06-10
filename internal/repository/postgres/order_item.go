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

// OrderItemRepositoryInterface define contract for orderItem related functions to repository
type OrderItemRepositoryInterface interface {
	GetOrderItemByID(ctx context.Context, orderItemID int) (*entity.OrderItem, error)
}

// OrderItemRepository holds database connection
type OrderItemRepository struct {
	db *sqlx.DB
}

var (
	// OrderItemTableName hold table name for order_items
	OrderItemTableName = "order_items"
	// OrderItemColumns list all columns on order_items table
	OrderItemColumns = []string{"id", "order_id", "book_id", "quantity", "price", "fee", "total_item_price", "created_at", "updated_at"}
	// OrderItemAttributes hold string format of all order_items table columns
	OrderItemAttributes = strings.Join(OrderItemColumns, ", ")
)

// NewOrderItemRepository create initiate orderItem repository with given database
func NewOrderItemRepository(db *sqlx.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.OrderItem, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.OrderItem, 0)

	for rows.Next() {
		tmpEntity := dbentity.OrderItem{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
}

// GetOrderItemByID query to get orderItem by ID
func (r *OrderItemRepository) GetOrderItemByID(ctx context.Context, orderItemID int) (*entity.OrderItem, error) {
	functionName := "OrderItemRepository.GetOrderItemByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1 LIMIT 1", OrderItemAttributes, OrderItemTableName)
	rows, err := r.fetch(ctx, query, orderItemID)
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if len(rows) == 0 {
		return nil, response.ErrNotFound
	}

	return rows[0], nil
}
