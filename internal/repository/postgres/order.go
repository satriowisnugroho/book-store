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

// OrderRepositoryInterface define contract for order related functions to repository
type OrderRepositoryInterface interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetOrdersByUserID(ctx context.Context, userID, limit, offset int) ([]*entity.Order, error)
	GetOrdersByUserIDCount(ctx context.Context, userID int) (int, error)
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

	// OrderCreationColumns list all columns used for create order
	OrderCreationColumns = OrderColumns[1:]
	// OrderCreationAttributes hold string format of all creation order columns
	OrderCreationAttributes = strings.Join(OrderCreationColumns, ", ")
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

// CreateOrder insert order data into database
func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	functionName := "OrderRepository.CreateOrder"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	now := time.Now()
	order.CreatedAt = now
	order.UpdatedAt = now

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, OrderTableName, OrderCreationAttributes, EnumeratedBindvars(OrderCreationColumns))

	err := r.db.QueryRowContext(ctx, query,
		order.UserID,
		order.BookID,
		order.Quantity,
		order.Price,
		order.Fee,
		order.TotalPrice,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(&order.ID)
	if err != nil {
		return errors.Wrap(err, functionName)
	}

	return nil
}

// GetOrdersByUserID query to get list of orders by user ID
func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID, limit, offset int) ([]*entity.Order, error) {
	functionName := "OrderRepository.GetOrdersByUserID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return []*entity.Order{}, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = %d LIMIT %d OFFSET %d", OrderAttributes, OrderTableName, userID, limit, offset)

	rows, err := r.fetch(ctx, query)
	if err != nil {
		return rows, errors.Wrap(err, functionName)
	}

	return rows, nil
}

// GetOrdersByUserIDCount query to get the count of orders by user ID
func (r *OrderRepository) GetOrdersByUserIDCount(ctx context.Context, userID int) (int, error) {
	functionName := "OrderRepository.GetOrdersByUserIDCount"
	if err := helper.CheckDeadline(ctx); err != nil {
		return 0, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = %d", OrderTableName, userID)

	count := 0
	rows := r.db.QueryRowxContext(ctx, query)
	if err := rows.Scan(&count); err != nil {
		return count, errors.Wrap(err, functionName)
	}

	return count, nil
}
