package usecase_test

import (
	"context"
	"errors"
	"testing"

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
		name      string
		ctx       context.Context
		payload   *entity.OrderPayload
		rBookRes  *entity.Book
		rBookErr  error
		rOrderErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:    "invalid payload",
			ctx:     context.Background(),
			payload: &entity.OrderPayload{Quantity: 0},
			wantErr: true,
		},
		{
			name:     "book is not found",
			ctx:      context.Background(),
			payload:  &entity.OrderPayload{Quantity: 1},
			rBookErr: response.ErrNotFound,
			wantErr:  true,
		},
		{
			name:     "failed to get book",
			ctx:      context.Background(),
			payload:  &entity.OrderPayload{Quantity: 1},
			rBookErr: errors.New("error get book"),
			wantErr:  true,
		},
		{
			name:      "failed to create order",
			ctx:       context.Background(),
			payload:   &entity.OrderPayload{Quantity: 1},
			rOrderErr: errors.New("error create order"),
			rBookRes:  &entity.Book{},
			wantErr:   true,
		},
		{
			name:     "success",
			ctx:      context.Background(),
			payload:  &entity.OrderPayload{Quantity: 1},
			rBookRes: &entity.Book{},
			wantErr:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			bookRepo := &testmock.BookRepositoryInterface{}
			bookRepo.On("GetBookByID", mock.Anything, mock.Anything).Return(tc.rBookRes, tc.rBookErr)

			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("CreateOrder", mock.Anything, mock.Anything).Return(tc.rOrderErr)

			uc := usecase.NewOrderUsecase(bookRepo, orderRepo)
			_, err := uc.CreateOrder(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestGetOrdersByUserID(t *testing.T) {
	testcases := []struct {
		name                  string
		ctx                   context.Context
		rGetOrdersByUserIDRes []*entity.Order
		rGetOrdersByUserIDErr error
		wantErr               bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:                  "failed to get orders",
			ctx:                   context.Background(),
			rGetOrdersByUserIDErr: errors.New("error get orders by user id"),
			wantErr:               true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("GetOrdersByUserID", mock.Anything, mock.Anything).Return(tc.rGetOrdersByUserIDRes, tc.rGetOrdersByUserIDErr)

			uc := usecase.NewOrderUsecase(&testmock.BookRepositoryInterface{}, orderRepo)
			_, err := uc.GetOrdersByUserID(tc.ctx)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
