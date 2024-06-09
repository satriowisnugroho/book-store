package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetLimitOffsetFromURLQuery get limit and offset from gin context
func GetLimitOffsetFromURLQuery(c *gin.Context) (int, int) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit > 20 || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset
}
