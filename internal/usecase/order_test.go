package usecase_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
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
		ctx       *gin.Context
		payload   *entity.OrderPayload
		rBookRes  *entity.Book
		rBookErr  error
		rOrderErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.GinCtxEnded(),
			wantErr: true,
		},
		{
			name:    "invalid payload",
			ctx:     fixture.GinCtxBackground(),
			payload: &entity.OrderPayload{Quantity: 0},
			wantErr: true,
		},
		{
			name:     "book is not found",
			ctx:      fixture.GinCtxBackground(),
			payload:  &entity.OrderPayload{Quantity: 1},
			rBookErr: response.ErrNotFound,
			wantErr:  true,
		},
		{
			name:     "failed to get book",
			ctx:      fixture.GinCtxBackground(),
			payload:  &entity.OrderPayload{Quantity: 1},
			rBookErr: errors.New("error get book"),
			wantErr:  true,
		},
		{
			name:      "failed to create order",
			ctx:       fixture.GinCtxBackground(),
			payload:   &entity.OrderPayload{Quantity: 1},
			rOrderErr: errors.New("error create order"),
			rBookRes:  &entity.Book{},
			wantErr:   true,
		},
		{
			name:     "success",
			ctx:      fixture.GinCtxBackground(),
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
		name                       string
		ctx                        *gin.Context
		rGetOrdersByUserIDRes      []*entity.Order
		rGetOrdersByUserIDErr      error
		rGetOrdersByUserIDCountRes int
		rGetOrdersByUserIDCountErr error
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
			name:                  "failed to get book",
			ctx:                   fixture.GinCtxBackground(),
			rGetOrdersByUserIDRes: []*entity.Order{{}},
			rGetBookByIDErr:       errors.New("error get book by id"),
			wantErr:               true,
		},
		{
			name:                  "success",
			ctx:                   fixture.GinCtxBackground(),
			rGetOrdersByUserIDRes: []*entity.Order{{}},
			wantErr:               false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			bookRepo := &testmock.BookRepositoryInterface{}
			bookRepo.On("GetBookByID", mock.Anything, mock.Anything).Return(tc.rGetBookByIDRes, tc.rGetBookByIDErr)

			orderRepo := &testmock.OrderRepositoryInterface{}
			orderRepo.On("GetOrdersByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.rGetOrdersByUserIDRes, tc.rGetOrdersByUserIDErr)
			orderRepo.On("GetOrdersByUserIDCount", mock.Anything, mock.Anything).Return(tc.rGetOrdersByUserIDCountRes, tc.rGetOrdersByUserIDCountErr)

			uc := usecase.NewOrderUsecase(bookRepo, orderRepo)
			_, _, err := uc.GetOrdersByUserID(tc.ctx, 10, 0)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
