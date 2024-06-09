package helper

import (
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext get the user ID from context
func GetUserIDFromContext(c *gin.Context) int {
	iUserID, _ := c.Get("user_id")
	userID, _ := iUserID.(int)

	return userID
}
