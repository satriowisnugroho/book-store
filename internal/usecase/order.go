package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// OrderUsecaseInterface define contract for order related functions to usecase
type OrderUsecaseInterface interface {
	CreateOrder(ctx context.Context, payload *entity.OrderPayload) (*entity.Order, error)
	GetOrdersByUserID(ctx context.Context) ([]*entity.Order, error)
}

type OrderUsecase struct {
	bookRepo  repo.BookRepositoryInterface
	orderRepo repo.OrderRepositoryInterface
}

func NewOrderUsecase(br repo.BookRepositoryInterface, or repo.OrderRepositoryInterface) *OrderUsecase {
	return &OrderUsecase{
		bookRepo:  br,
		orderRepo: or,
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

	book, err := uc.bookRepo.GetBookByID(ctx, payload.BookID)
	if err != nil {
		if err == response.ErrNotFound {
			return nil, err
		}

		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetBookByID: %w", err), functionName)
	}

	order := &entity.Order{}
	// TODO: Get from current user
	order.UserID = 1
	order.BookID = book.ID
	order.Quantity = payload.Quantity
	order.Price = book.Price
	order.Fee = config.ServiceFee
	order.TotalPrice = (payload.Quantity * book.Price) + config.ServiceFee
	if err := uc.orderRepo.CreateOrder(ctx, order); err != nil {
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
	orders, err := uc.orderRepo.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetOrdersByUserID: %w", err), functionName)
	}

	return orders, nil
}
