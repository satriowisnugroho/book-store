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

func (uc *OrderUsecase) GetOrdersByUserID(ctx context.Context) ([]*entity.Order, error) {
	functionName := "OrderUsecase.GetOrdersByUserID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	// TODO: Get from current user
	userID := int64(1)
	orders, err := uc.repo.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetOrdersByUserID: %w", err), functionName)
	}

	return orders, nil
}
