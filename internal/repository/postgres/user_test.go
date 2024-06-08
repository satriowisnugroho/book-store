package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/test/fixture"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		input     *entity.User
		createErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:      "duplicate sku & tenant",
			ctx:       context.Background(),
			input:     &entity.User{},
			createErr: &pq.Error{Code: pq.ErrorCode(config.UniqueConstraintViolationCode)},
			wantErr:   true,
		},
		{
			name:      "fail exec query",
			ctx:       context.Background(),
			input:     &entity.User{},
			createErr: errors.New("fail exec"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			input:   &entity.User{},
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

			expectedQuery := "INSERT INTO .+ (.+) VALUES (.+) RETURNING id"
			if tc.createErr != nil {
				mock.ExpectQuery(expectedQuery).WillReturnError(tc.createErr)
			} else {
				row := sqlmock.NewRows([]string{"id"})
				result := row.AddRow(1)
				mock.ExpectQuery(expectedQuery).WillReturnRows(result)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewUserRepository(dbx)

			err = repo.CreateUser(tc.ctx, tc.input)
			assert.Equal(t, tc.wantErr, err != nil)
			if !tc.wantErr {
				assert.Equal(t, 1, tc.input.ID)
			}
		})
	}
}
