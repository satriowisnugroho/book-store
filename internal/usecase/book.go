package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
)

// BookUsecaseInterface define contract for book related functions to usecase
type BookUsecaseInterface interface {
	GetBooks(ctx context.Context, payload entity.GetBooksPayload) ([]*entity.Book, int, error)
}

type BookUsecase struct {
	repo repo.BookRepositoryInterface
}

func NewBookUsecase(r repo.BookRepositoryInterface) *BookUsecase {
	return &BookUsecase{
		repo: r,
	}
}

func (uc *BookUsecase) GetBooks(ctx context.Context, payload entity.GetBooksPayload) ([]*entity.Book, int, error) {
	functionName := "BookUsecase.GetBooks"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, 0, errors.Wrap(err, functionName)
	}

	books, err := uc.repo.GetBooks(ctx, payload)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.repo.GetBooks: %w", err), functionName)
	}

	count, err := uc.repo.GetBooksCount(ctx, payload)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.repo.GetBooksCount: %w", err), functionName)
	}

	return books, count, nil
}
