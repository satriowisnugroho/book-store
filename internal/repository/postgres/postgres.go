package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
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

// UpdateColumnsValues is func to convert list columns to update query
func UpdateColumnsValues(columns []string) string {
	var keyValues []string
	for i, column := range columns {
		keyValues = append(keyValues, fmt.Sprintf("%s = $%d", column, i+1))
	}

	return strings.Join(keyValues, ", ")
}

// Tx get db or transaction
func Tx(db *sqlx.DB, iTrx interface{}) sqlx.ExtContext {
	if iTrx == nil {
		return db
	}

	return iTrx.(*sqlx.Tx)
}

func isUniqueConstraintViolation(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && pqErr.Code == config.UniqueConstraintViolationCode
}
