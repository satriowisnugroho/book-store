package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/satriowisnugroho/book-store/internal/handler/http/middleware"
	"github.com/stretchr/testify/assert"
)

func setupTestRouterAuthMiddleware(secret string) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware(secret))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized!"})
	})

	return r
}

func TestAuthMiddleware(t *testing.T) {
	secret := "secret"
	router := setupTestRouterAuthMiddleware(secret)

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "no token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid format",
			token:          "invalid-format",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid token",
			token:          "Bearer invalid-format",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "valid token",
			token: func() string {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"foo": "bar"})
				tokenString, _ := token.SignedString([]byte(secret))
				return "Bearer " + tokenString
			}(),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/protected", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
