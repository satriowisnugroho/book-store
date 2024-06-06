package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
)

// OrderUsecaseInterface define contract for order related functions to usecase
type OrderUsecaseInterface interface {
	CreateOrder(ctx context.Context, payload *entity.OrderPayload) (*entity.Order, error)
	GetOrdersByUserID(ctx context.Context) ([]*entity.Order, error)
}

type OrderUsecase struct {
	repo repo.OrderRepositoryInterface
}

func NewOrderUsecase(r repo.OrderRepositoryInterface) *OrderUsecase {
	return &OrderUsecase{
		repo: r,
	}
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, payload *entity.OrderPayload) (*entity.Order, error) {
	functionName := "OrderUsecase.CreateOrder"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	// TODO: Get book by ID

	order := &entity.Order{}
	order.UserID = 1
	order.BookID = payload.BookID
	order.Quantity = payload.Quantity
	order.Price = 0
	order.Fee = 0
	order.TotalPrice = (payload.Quantity * 0) + 0

	if err := uc.repo.CreateOrder(ctx, order); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateOrder: %w", err), functionName)
	}

	return order, nil
}

func (uc *OrderUsecase) GetOrdersByUserID(ctx context.Context) ([]*entity.Order, error) {
	functionName := "OrderUsecase.GetOrdersByUserID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	// TODO: Get from current user
	userID := 1
	orders, err := uc.repo.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetOrdersByUserID: %w", err), functionName)
	}

	return orders, nil
}
