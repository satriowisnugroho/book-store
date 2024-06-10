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

func TestCreateOrderItem(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		input     *entity.OrderItem
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
			input:     &entity.OrderItem{},
			createErr: errors.New("fail exec"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			input:   &entity.OrderItem{},
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
			repo := postgres.NewOrderRepository(dbx)

			err = repo.CreateOrderItem(tc.ctx, tc.input)
			assert.Equal(t, tc.wantErr, err != nil)
			if !tc.wantErr {
				assert.Equal(t, 1, tc.input.ID)
			}
		})
	}
}

func TestGetOrderItemByID(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		fetchErr  error
		fetchRows []string
		expected  *entity.OrderItem
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
			name:      "record not found",
			ctx:       context.Background(),
			fetchRows: postgres.OrderItemColumns,
			wantErr:   true,
		},
		{
			name:      "success",
			ctx:       context.Background(),
			fetchRows: postgres.OrderItemColumns,
			expected:  &entity.OrderItem{},
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

			mockExpectedQuery := mock.ExpectQuery("^SELECT .+ FROM .+ WHERE id = .+ LIMIT 1")
			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected.ID,
						tc.expected.OrderID,
						tc.expected.BookID,
						tc.expected.Quantity,
						tc.expected.Price,
						tc.expected.Fee,
						tc.expected.TotalItemPrice,
						tc.expected.CreatedAt,
						tc.expected.UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewOrderItemRepository(dbx)
			result, err := repo.GetOrderItemByID(tc.ctx, 123)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}
