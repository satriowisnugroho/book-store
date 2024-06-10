package usecase_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/test/fixture"
	testmock "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	testcases := []struct {
		name                string
		ctx                 *gin.Context
		payload             *entity.OrderPayload
		rStartTrxErr        error
		rCommitTrxErr       error
		rBookRes            *entity.Book
		rBookErr            error
		rCreateOrderErr     error
		rUpdateOrderErr     error
		rCreateOrderItemErr error
		wantErr             bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.GinCtxEnded(),
			wantErr: true,
		},
		{
			name:    "invalid payload",
			ctx:     fixture.GinCtxBackground(),
			payload: &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 0}}},
			wantErr: true,
		},
		{
			name:         "failed to start transaction",
			ctx:          fixture.GinCtxBackground(),
			payload:      &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rStartTrxErr: response.ErrNoSQLTransactionFound,
			wantErr:      true,
		},
		{
			name:            "failed to create order",
			ctx:             fixture.GinCtxBackground(),
			payload:         &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rCreateOrderErr: errors.New("error create order"),
			wantErr:         true,
		},
		{
			name:     "book is not found",
			ctx:      fixture.GinCtxBackground(),
			payload:  &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rBookErr: response.ErrNotFound,
			wantErr:  true,
		},
		{
			name:     "failed to get book",
			ctx:      fixture.GinCtxBackground(),
			payload:  &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rBookErr: errors.New("error get book"),
			wantErr:  true,
		},
		{
			name:            "failed to update order",
			ctx:             fixture.GinCtxBackground(),
			payload:         &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rBookRes:        &entity.Book{},
			rUpdateOrderErr: errors.New("error update order"),
			wantErr:         true,
		},
		{
			name:          "failed to commit transaction",
			ctx:           fixture.GinCtxBackground(),
			payload:       &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rCommitTrxErr: response.ErrNoSQLTransactionFound,
			rBookRes:      &entity.Book{},
			wantErr:       true,
		},
		{
			name:     "success",
			ctx:      fixture.GinCtxBackground(),
			payload:  &entity.OrderPayload{OrderItems: []entity.OrderItemPayload{{Quantity: 1}}},
			rBookRes: &entity.Book{},
			wantErr:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			dbTransactionRepo := &testmock.PostgresTransactionRepositoryInterface{}
			dbTransactionRepo.On("StartTransactionQuery", mock.Anything).Return(&sqlx.Tx{}, tc.rStartTrxErr)
			dbTransactionRepo.On("CommitTransactionQuery", mock.Anything, mock.Anything).Return(tc.rCommitTrxErr)
			dbTransactionRepo.On("RollbackTransactionQuery", mock.Anything, mock.Anything).Return(nil)

			bookRepo := &testmock.BookRepositoryInterface{}
			bookRepo.On("GetBookByID", mock.Anything, mock.Anything).Return(tc.rBookRes, tc.rBookErr)

			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything).Return(tc.rCreateOrderErr)
			orderRepo.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(tc.rUpdateOrderErr)

			orderItemRepo := &testmock.OrderItemRepositoryInterface{}
			orderItemRepo.On("CreateOrderItem", mock.Anything, mock.Anything, mock.Anything).Return(tc.rCreateOrderItemErr)

			uc := usecase.NewOrderUsecase(dbTransactionRepo, bookRepo, orderRepo, orderItemRepo)
			_, err := uc.CreateOrder(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestGetOrdersByUserID(t *testing.T) {
	testcases := []struct {
		name                       string
		ctx                        *gin.Context
		rGetOrdersByUserIDRes      []*entity.Order
		rGetOrdersByUserIDErr      error
		rGetOrdersByUserIDCountRes int
		rGetOrdersByUserIDCountErr error
		rGetOrderItemsByOrderIDRes []*entity.OrderItem
		rGetOrderItemsByOrderIDErr error
		rGetBookByIDRes            *entity.Book
		rGetBookByIDErr            error
		wantErr                    bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.GinCtxEnded(),
			wantErr: true,
		},
		{
			name:                  "failed to get orders",
			ctx:                   fixture.GinCtxBackground(),
			rGetOrdersByUserIDErr: errors.New("error get orders by user id"),
			wantErr:               true,
		},
		{
			name:                       "failed to get orders count",
			ctx:                        fixture.GinCtxBackground(),
			rGetOrdersByUserIDCountErr: errors.New("error get orders by user id count"),
			wantErr:                    true,
		},
		{
			name:                       "failed to get order items",
			ctx:                        fixture.GinCtxBackground(),
			rGetOrdersByUserIDRes:      []*entity.Order{{}},
			rGetOrderItemsByOrderIDErr: errors.New("error get order items by order id"),
			wantErr:                    true,
		},
		{
			name:                       "failed to get book",
			ctx:                        fixture.GinCtxBackground(),
			rGetOrdersByUserIDRes:      []*entity.Order{{}},
			rGetOrderItemsByOrderIDRes: []*entity.OrderItem{{}},
			rGetBookByIDErr:            errors.New("error get book by id"),
			wantErr:                    true,
		},
		{
			name:                       "success",
			ctx:                        fixture.GinCtxBackground(),
			rGetOrdersByUserIDRes:      []*entity.Order{{}},
			rGetOrderItemsByOrderIDRes: []*entity.OrderItem{{}},
			wantErr:                    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			bookRepo := &testmock.BookRepositoryInterface{}
			bookRepo.On("GetBookByID", mock.Anything, mock.Anything).Return(tc.rGetBookByIDRes, tc.rGetBookByIDErr)

			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("GetOrdersByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.rGetOrdersByUserIDRes, tc.rGetOrdersByUserIDErr)
			orderRepo.On("GetOrdersByUserIDCount", mock.Anything, mock.Anything).Return(tc.rGetOrdersByUserIDCountRes, tc.rGetOrdersByUserIDCountErr)

			orderItemRepo := &testmock.OrderItemRepositoryInterface{}
			orderItemRepo.On("GetOrderItemsByOrderID", mock.Anything, mock.Anything).Return(tc.rGetOrderItemsByOrderIDRes, tc.rGetOrderItemsByOrderIDErr)

			uc := usecase.NewOrderUsecase(&testmock.PostgresTransactionRepositoryInterface{}, bookRepo, orderRepo, orderItemRepo)
			_, _, err := uc.GetOrdersByUserID(tc.ctx, 10, 0)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
