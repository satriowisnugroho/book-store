package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/test/fixture"
	testmock "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBooks(t *testing.T) {
	testcases := []struct {
		name         string
		ctx          context.Context
		rGetBooksRes []*entity.Book
		rGetBooksErr error
		wantErr      bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:         "failed to get books",
			ctx:          context.Background(),
			rGetBooksErr: errors.New("error get books"),
			wantErr:      true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			bookRepo := &testmock.BookRepositoryInterface{}
			bookRepo.On("GetBooks", mock.Anything).Return(tc.rGetBooksRes, tc.rGetBooksErr)

			uc := usecase.NewBookUsecase(bookRepo)
			_, err := uc.GetBooks(tc.ctx)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
