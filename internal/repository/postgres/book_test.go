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

func TestGetBooks(t *testing.T) {
	testcases := []struct {
		name        string
		ctx         context.Context
		fetchErr    error
		fetchRows   []string
		payload     entity.GetBooksPayload
		filterQuery string
		expected    []*entity.Book
		wantErr     bool
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
			name:        "success",
			ctx:         context.Background(),
			fetchRows:   postgres.BookColumns,
			payload:     entity.GetBooksPayload{TitleKeyword: "foo"},
			filterQuery: "title ILIKE \\$1",
			expected:    []*entity.Book{{}},
			wantErr:     false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			expectedQuery := "SELECT .+ FROM books"
			if tc.filterQuery != "" {
				expectedQuery = expectedQuery + " WHERE " + tc.filterQuery
			}
			expectedQuery = expectedQuery + " LIMIT .+ OFFSET .+"
			mockExpectedQuery := mock.ExpectQuery(expectedQuery)

			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected[0].ID,
						tc.expected[0].Isbn,
						tc.expected[0].Title,
						tc.expected[0].Price,
						tc.expected[0].CreatedAt,
						tc.expected[0].UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewBookRepository(dbx)
			result, err := repo.GetBooks(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestGetBooksCount(t *testing.T) {
	testcases := []struct {
		name        string
		ctx         context.Context
		fetchErr    error
		payload     entity.GetBooksPayload
		filterQuery string
		expected    int
		wantErr     bool
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
			name:        "success",
			ctx:         context.Background(),
			payload:     entity.GetBooksPayload{TitleKeyword: "foo"},
			filterQuery: "title ILIKE \\$1",
			expected:    1,
			wantErr:     false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			expectedQuery := "SELECT COUNT\\(\\*\\) FROM books"
			if tc.filterQuery != "" {
				expectedQuery = expectedQuery + " WHERE " + tc.filterQuery
			}
			mockExpectedQuery := mock.ExpectQuery(expectedQuery)

			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows([]string{"COUNT(*)"})
				rows = rows.AddRow(tc.expected)

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewBookRepository(dbx)
			result, err := repo.GetBooksCount(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestGetBookByID(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		fetchErr  error
		fetchRows []string
		expected  *entity.Book
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
			fetchRows: postgres.BookColumns,
			wantErr:   true,
		},
		{
			name:      "success",
			ctx:       context.Background(),
			fetchRows: postgres.BookColumns,
			expected:  &entity.Book{},
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

			mockExpectedQuery := mock.ExpectQuery("^SELECT .+ FROM books WHERE id = .+ LIMIT 1")
			if tc.fetchErr != nil {
				mockExpectedQuery.WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected.ID,
						tc.expected.Isbn,
						tc.expected.Title,
						tc.expected.Price,
						tc.expected.CreatedAt,
						tc.expected.UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mockExpectedQuery.WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewBookRepository(dbx)
			result, err := repo.GetBookByID(tc.ctx, 123)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}
