package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/response"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(config.AuthorizationHeader)
		if authHeader == "" {
			response.Error(c, response.ErrUnauthorized("Authorization header required"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != config.AuthorizationHeaderBearer {
			response.Error(c, response.ErrUnauthorized("Invalid Authorization header format"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, response.ErrUnauthorized("Invalid token"))
			c.Abort()
			return
		}

		c.Set("user_id", (*claims)["user_id"])
		c.Set("email", (*claims)["email"])
		c.Next()
	}
}
