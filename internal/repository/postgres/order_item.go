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
)

// OrderItemRepositoryInterface define contract for orderItem related functions to repository
type OrderItemRepositoryInterface interface {
	CreateOrderItem(ctx context.Context, dbTrx interface{}, orderItem *entity.OrderItem) error
	GetOrderItemsByOrderID(ctx context.Context, orderID int) ([]*entity.OrderItem, error)
}

// OrderItemRepository holds database connection
type OrderItemRepository struct {
	db *sqlx.DB
}

var (
	// OrderItemTableName hold table name for order_items
	OrderItemTableName = "order_items"
	// OrderItemColumns list all columns on order_items table
	OrderItemColumns = []string{"id", "order_id", "book_id", "quantity", "price", "total_item_price", "created_at", "updated_at"}
	// OrderItemAttributes hold string format of all order_items table columns
	OrderItemAttributes = strings.Join(OrderItemColumns, ", ")

	// OrderItemCreationColumns list all columns used for create orderItem
	OrderItemCreationColumns = OrderItemColumns[1:]
	// OrderItemCreationAttributes hold string format of all creation orderItem columns
	OrderItemCreationAttributes = strings.Join(OrderItemCreationColumns, ", ")
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

// CreateOrderItem insert orderItem data into database
func (r *OrderItemRepository) CreateOrderItem(ctx context.Context, dbTrx interface{}, orderItem *entity.OrderItem) error {
	functionName := "OrderItemRepository.CreateOrderItem"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	now := time.Now()
	orderItem.CreatedAt = now
	orderItem.UpdatedAt = now

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, OrderItemTableName, OrderItemCreationAttributes, EnumeratedBindvars(OrderItemCreationColumns))

	tx := Tx(r.db, dbTrx)
	err := tx.QueryRowxContext(
		ctx,
		query,
		orderItem.OrderID,
		orderItem.BookID,
		orderItem.Quantity,
		orderItem.Price,
		orderItem.TotalItemPrice,
		orderItem.CreatedAt,
		orderItem.UpdatedAt,
	).Scan(&orderItem.ID)
	if err != nil {
		return errors.Wrap(err, functionName)
	}

	return nil
}

// GetOrderItemsByOrderID query to get orderItem by order ID
func (r *OrderItemRepository) GetOrderItemsByOrderID(ctx context.Context, orderID int) ([]*entity.OrderItem, error) {
	functionName := "OrderItemRepository.GetOrderItemsByOrderID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE order_id = %d", OrderItemAttributes, OrderItemTableName, orderID)
	rows, err := r.fetch(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	return rows, nil
}
