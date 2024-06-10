package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/response"
)

// PostgresTransactionRepository holds database connection
type PostgresTransactionRepository struct {
	db *sqlx.DB
}

// PostgresTransactionRepositoryInterface define contract for postgres transaction related functions to repository
type PostgresTransactionRepositoryInterface interface {
	StartTransactionQuery(ctx context.Context) (*sqlx.Tx, error)
	CommitTransactionQuery(ctx context.Context, tx *sqlx.Tx) error
	RollbackTransactionQuery(ctx context.Context, tx *sqlx.Tx) error
}

// NewPostgresTransactionRepository create initiate postgres transaction repository with given database
func NewPostgresTransactionRepository(db *sqlx.DB) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

// StartTransactionQuery create transaction instance
func (r *PostgresTransactionRepository) StartTransactionQuery(ctx context.Context) (*sqlx.Tx, error) {
	functionName := "PostgresTransactionRepository.StartTransactionQuery"
	newTx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	return newTx, errors.Wrap(err, functionName)
}

// CommitTransactionQuery commit transaction instance
func (r *PostgresTransactionRepository) CommitTransactionQuery(ctx context.Context, tx *sqlx.Tx) error {
	functionName := "PostgresTransactionRepository.CommitTransactionQuery"
	if tx == nil {
		return response.ErrNoSQLTransactionFound
	}

	err := tx.Commit()
	return errors.Wrap(err, functionName)
}

// RollbackTransactionQuery rollback transaction instance
func (r *PostgresTransactionRepository) RollbackTransactionQuery(ctx context.Context, tx *sqlx.Tx) error {
	functionName := "PostgresTransactionRepository.RollbackTransactionQuery"
	if tx == nil {
		return response.ErrNoSQLTransactionFound
	}

	err := tx.Rollback()
	return errors.Wrap(err, functionName)
}
