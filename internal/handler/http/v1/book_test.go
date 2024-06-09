package v1_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/entity"
	httpv1 "github.com/satriowisnugroho/book-store/internal/handler/http/v1"
	testmock "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBooks(t *testing.T) {
	testcases := []struct {
		name              string
		uBookErr          error
		httpStatusCodeRes int
	}{
		{
			name:              "failed to get books",
			uBookErr:          errors.New("error get books"),
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

			ctx.Request, _ = http.NewRequest("GET", "/books", nil)

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			bookUsecase := &testmock.BookUsecaseInterface{}
			bookUsecase.On("GetBooks", mock.Anything, mock.Anything).Return([]*entity.Book{{}}, 10, tc.uBookErr)

			h := &httpv1.BookHandler{l, bookUsecase}
			h.GetBooks(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}
