package fixture

import (
	"context"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
)

// CtxEnded creates dummy context with cancelled state
func CtxEnded() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))
	defer cancel()
	return ctx
}

// GinCtxEnded creates dummy gin context with cancelled state
func GinCtxEnded() *gin.Context {
	c := GinCtx()
	req, _ := http.NewRequestWithContext(CtxEnded(), http.MethodGet, "/test", nil)
	c.Request = req
	return c
}

// GinCtxBackground creates dummy gin context with background context
func GinCtxBackground() *gin.Context {
	c := GinCtx()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/test", nil)
	c.Request = req
	return c
}

// GinCtx creates dummy gin context
func GinCtx() *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c
}
