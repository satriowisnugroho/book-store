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

func TestRegister(t *testing.T) {
	testcases := []struct {
		name              string
		body              string
		uUserRes          *entity.User
		uUserErr          error
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
			uUserErr:          errors.New("error create order"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			body:              `{}`,
			uUserRes:          &entity.User{},
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

			orderUsecase := &testmock.UserUsecaseInterface{}
			orderUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(tc.uUserRes, tc.uUserErr)

			h := &httpv1.UserHandler{l, orderUsecase}
			h.Register(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}
