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
	GetOrdersByUserID(c *gin.Context, limit, offset int) ([]*entity.OrderResponse, int, error)
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

	// Validate the payload
	for _, orderItemPayload := range payload.OrderItems {
		if err := orderItemPayload.Validate(); err != nil {
			return nil, err
		}
	}

	// TODO: Start trx

	order := &entity.Order{}
	order.UserID = helper.GetUserIDFromContext(c)
	order.Fee = config.ServiceFee
	order.TotalPrice = config.ServiceFee
	if err := uc.orderRepo.CreateOrder(ctx, order); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateOrder: %w", err), functionName)
	}

	totalPrice := 0
	for _, orderItemPayload := range payload.OrderItems {
		book, err := uc.bookRepo.GetBookByID(ctx, orderItemPayload.BookID)
		if err != nil {
			if err == response.ErrNotFound {
				return nil, err
			}

			return nil, errors.Wrap(fmt.Errorf("uc.repo.GetBookByID: %w", err), functionName)
		}

		orderItem := &entity.OrderItem{}
		orderItem.OrderID = order.ID
		orderItem.BookID = book.ID
		orderItem.Quantity = orderItemPayload.Quantity
		orderItem.Price = book.Price
		orderItem.TotalItemPrice = orderItemPayload.Quantity * book.Price
		totalPrice += orderItem.TotalItemPrice
	}

	order.TotalPrice += totalPrice
	// if err := uc.orderRepo.UpdateOrder(ctx, order); err != nil {
	// 	return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateOrder: %w", err), functionName)
	// }

	// TODO: commit trx

	return order, nil
}

func (uc *OrderUsecase) GetOrdersByUserID(c *gin.Context, limit, offset int) ([]*entity.OrderResponse, int, error) {
	functionName := "OrderUsecase.GetOrdersByUserID"

	ctx := c.Request.Context()
	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, 0, errors.Wrap(err, functionName)
	}

	userID := helper.GetUserIDFromContext(c)
	orders, err := uc.orderRepo.GetOrdersByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.orderRepo.GetOrdersByUserID: %w", err), functionName)
	}

	count, err := uc.orderRepo.GetOrdersByUserIDCount(ctx, userID)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.orderRepo.GetOrdersByUserIDCount: %w", err), functionName)
	}

	orderRes := []*entity.OrderResponse{}
	for _, order := range orders {
		// book, err := uc.bookRepo.GetBookByID(ctx, order.BookID)
		// if err != nil {
		// 	return nil, 0, errors.Wrap(fmt.Errorf("uc.bookRepo.GetBookByID: %w", err), functionName)
		// }

		orderRes = append(orderRes, &entity.OrderResponse{Order: order})
	}

	return orderRes, count, nil
}
