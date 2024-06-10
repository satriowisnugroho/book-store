package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
)

func TestStartTransactionQuery(t *testing.T) {
	testcases := []struct {
		name     string
		ctx      context.Context
		errBegin error
		wantErr  bool
	}{
		{
			name:     "error begin tx",
			ctx:      context.Background(),
			errBegin: errors.New("error begin tx"),
			wantErr:  true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectBegin().WillReturnError(tc.errBegin)

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewPostgresTransactionRepository(dbx)

			tx, err := repo.StartTransactionQuery(tc.ctx)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, tx)
			}
		})
	}
}

func TestCommitTransactionQuery(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		errCommit error
		txNil     bool
		wantErr   bool
	}{
		{
			name:    "error tx nil",
			ctx:     context.Background(),
			txNil:   true,
			wantErr: true,
		},
		{
			name:      "error commit tx",
			ctx:       context.Background(),
			errCommit: errors.New("error commit tx"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var tx *sqlx.Tx
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectBegin().WillReturnError(nil)
			mock.ExpectCommit().WillReturnError(tc.errCommit)

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewPostgresTransactionRepository(dbx)

			if !tc.txNil {
				tx, _ = repo.StartTransactionQuery(tc.ctx)
			}

			err = repo.CommitTransactionQuery(tc.ctx, tx)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRollbackTransactionQuery(t *testing.T) {
	testcases := []struct {
		name        string
		ctx         context.Context
		errRollback error
		txNil       bool
		wantErr     bool
	}{
		{
			name:    "error tx nil",
			ctx:     context.Background(),
			txNil:   true,
			wantErr: true,
		},
		{
			name:        "error rollback tx",
			ctx:         context.Background(),
			errRollback: errors.New("error rollback tx"),
			wantErr:     true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var tx *sqlx.Tx
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectBegin().WillReturnError(nil)
			mock.ExpectRollback().WillReturnError(tc.errRollback)

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewPostgresTransactionRepository(dbx)

			if !tc.txNil {
				tx, _ = repo.StartTransactionQuery(tc.ctx)
			}

			err = repo.RollbackTransactionQuery(tc.ctx, tx)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
