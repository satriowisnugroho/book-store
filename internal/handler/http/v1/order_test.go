package v1_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/entity"
	httpv1 "github.com/satriowisnugroho/book-store/internal/handler/http/v1"
	testmock "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	testcases := []struct {
		name              string
		body              string
		uOrderRes         *entity.Order
		uOrderErr         error
		httpStatusCodeRes int
	}{
		{
			name:              "failed to decode payload",
			body:              `{failed}`,
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "failed to create order",
			body:              `{}`,
			uOrderErr:         errors.New("error create order"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			body:              `{}`,
			uOrderRes:         &entity.Order{},
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
				Method: "POST",
				Body:   io.NopCloser(strings.NewReader(tc.body)),
			}

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			orderUsecase := &testmock.OrderUsecaseInterface{}
			orderUsecase.On("CreateOrder", mock.Anything, mock.Anything).Return(tc.uOrderRes, tc.uOrderErr)

			h := &httpv1.OrderHandler{l, orderUsecase}
			h.CreateOrder(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}

func TestGetOrderHistory(t *testing.T) {
	testcases := []struct {
		name              string
		uOrderErr         error
		httpStatusCodeRes int
	}{
		{
			name:              "failed to get history of orders",
			uOrderErr:         errors.New("error get history of orders"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request, _ = http.NewRequest("GET", "/orders", nil)

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			orderUsecase := &testmock.OrderUsecaseInterface{}
			orderUsecase.On("GetOrdersByUserID", mock.Anything).Return([]*entity.OrderResponse{{}}, tc.uOrderErr)

			h := &httpv1.OrderHandler{l, orderUsecase}
			h.GetOrderHistory(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}
