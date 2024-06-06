package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/satriowisnugroho/book-store/internal/entity"
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
			name:      "failed to create order",
			ctx:       context.Background(),
			payload:   &entity.OrderPayload{Quantity: 1},
			rOrderErr: errors.New("error create order"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			payload: &entity.OrderPayload{Quantity: 1},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("CreateOrder", mock.Anything, mock.Anything).Return(tc.rOrderErr)

			uc := usecase.NewOrderUsecase(orderRepo)
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

			uc := usecase.NewOrderUsecase(orderRepo)
			_, err := uc.GetOrdersByUserID(tc.ctx)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
