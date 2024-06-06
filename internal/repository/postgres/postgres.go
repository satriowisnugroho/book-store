package postgres

import (
	"fmt"
	"strings"
)

// EnumeratedBindvars is func to convert list columns to bindvars
func EnumeratedBindvars(columns []string) string {
	var values []string
	for i := range columns {
		values = append(values, fmt.Sprintf("$%d", i+1))
	}

	return strings.Join(values, ", ")
}
