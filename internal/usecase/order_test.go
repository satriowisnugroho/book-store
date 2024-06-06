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
