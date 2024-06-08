package postgres

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/satriowisnugroho/book-store/internal/config"
)

// EnumeratedBindvars is func to convert list columns to bindvars
func EnumeratedBindvars(columns []string) string {
	var values []string
	for i := range columns {
		values = append(values, fmt.Sprintf("$%d", i+1))
	}

	return strings.Join(values, ", ")
}

func isUniqueConstraintViolation(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && pqErr.Code == config.UniqueConstraintViolationCode
}
