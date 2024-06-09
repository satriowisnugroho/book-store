package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// OrderUsecaseInterface define contract for order related functions to usecase
type OrderUsecaseInterface interface {
	CreateOrder(c *gin.Context, payload *entity.OrderPayload) (*entity.Order, error)
	GetOrdersByUserID(c *gin.Context) ([]*entity.Order, error)
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

func (uc *OrderUsecase) CreateOrder(c *gin.Context, payload *entity.OrderPayload) (*entity.Order, error) {
	functionName := "OrderUsecase.CreateOrder"

	ctx := c.Request.Context()
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
	order.UserID = helper.GetUserIDFromContext(c)
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

func (uc *OrderUsecase) GetOrdersByUserID(c *gin.Context) ([]*entity.Order, error) {
	functionName := "OrderUsecase.GetOrdersByUserID"

	ctx := c.Request.Context()
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	orders, err := uc.orderRepo.GetOrdersByUserID(ctx, helper.GetUserIDFromContext(c))
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetOrdersByUserID: %w", err), functionName)
	}

	return orders, nil
}
