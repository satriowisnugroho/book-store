package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/config"
	v1 "github.com/satriowisnugroho/book-store/internal/handler/http/v1"
	mocks "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	r := gin.Default()
	v1.NewRouter(r, &mocks.LoggerInterface{}, &config.Config{}, &mocks.BookUsecaseInterface{}, &mocks.OrderUsecaseInterface{}, &mocks.UserUsecaseInterface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
