package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/test/fixture"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		input     *entity.Order
		createErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:      "fail exec query",
			ctx:       context.Background(),
			input:     &entity.Order{},
			createErr: errors.New("fail exec"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			input:   &entity.Order{},
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

			expectedQuery := "INSERT INTO orders (.+) VALUES (.+) RETURNING id"
			if tc.createErr != nil {
				mock.ExpectQuery(expectedQuery).WillReturnError(tc.createErr)
			} else {
				row := sqlmock.NewRows([]string{"id"})
				result := row.AddRow(1)
				mock.ExpectQuery(expectedQuery).WillReturnRows(result)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewOrderRepository(dbx)

			err = repo.CreateOrder(tc.ctx, nil, tc.input)
			assert.Equal(t, tc.wantErr, err != nil)
			if !tc.wantErr {
				assert.Equal(t, 1, tc.input.ID)
			}
		})
	}
}

func TestGetOrdersByUserID(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		fetchErr  error
		fetchRows []string
		expected  []*entity.Order
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:     "fail fetch query error",
			ctx:      context.Background(),
			fetchErr: errors.New("fail fetch"),
			wantErr:  true,
		},
		{
			name:      "fail fetch return error rows",
			ctx:       context.Background(),
			fetchRows: []string{"unknown_column"},
			wantErr:   true,
		},
		{
			name:      "success",
			ctx:       context.Background(),
			fetchRows: postgres.OrderColumns,
			expected:  []*entity.Order{{}},
			wantErr:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockExpectedQuery := mock.ExpectQuery("SELECT .+ FROM orders WHERE user_id = .+ ORDER BY .+ LIMIT .+ OFFSET .+")
			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected[0].ID,
						tc.expected[0].UserID,
						tc.expected[0].Fee,
						tc.expected[0].TotalPrice,
						tc.expected[0].CreatedAt,
						tc.expected[0].UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewOrderRepository(dbx)
			result, err := repo.GetOrdersByUserID(tc.ctx, 1, 10, 0)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestGetOrdersByUserIDCount(t *testing.T) {
	testcases := []struct {
		name     string
		ctx      context.Context
		fetchErr error
		expected int
		wantErr  bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:     "fail fetch query error",
			ctx:      context.Background(),
			fetchErr: errors.New("fail fetch"),
			wantErr:  true,
		},
		{
			name:     "success",
			ctx:      context.Background(),
			expected: 1,
			wantErr:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockExpectedQuery := mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM orders WHERE user_id = .+")
			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows([]string{"COUNT(*)"})
				rows = rows.AddRow(tc.expected)

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewOrderRepository(dbx)
			result, err := repo.GetOrdersByUserIDCount(tc.ctx, 1)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		updateErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:      "fail exec query",
			ctx:       context.Background(),
			updateErr: errors.New("fail exec"),
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
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockExpectedQuery := mock.ExpectExec("UPDATE orders SET .+ WHERE id = .+")
			if tc.updateErr != nil {
				mockExpectedQuery.WillReturnError(tc.updateErr)
			} else {
				mockExpectedQuery.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewOrderRepository(dbx)
			err = repo.UpdateOrder(tc.ctx, nil, &entity.Order{})
			assert.Equal(t, tc.wantErr, err != nil, err)
		})
	}
}
