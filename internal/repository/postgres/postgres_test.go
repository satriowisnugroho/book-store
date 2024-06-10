package postgres_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
)

func TestTx(t *testing.T) {
	dbTrx := &sqlx.Tx{}
	tx := postgres.Tx(nil, &sqlx.Tx{})

	assert.Equal(t, dbTrx, tx)
}
