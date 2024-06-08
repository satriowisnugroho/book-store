package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecaseInterface define contract for user related functions to usecase
type UserUsecaseInterface interface {
	CreateUser(ctx context.Context, payload *entity.UserPayload) (*entity.User, error)
}

type UserUsecase struct {
	userRepo repo.UserRepositoryInterface
}

func NewUserUsecase(ur repo.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, payload *entity.UserPayload) (*entity.User, error) {
	functionName := "UserUsecase.CreateUser"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	// TODO: Move to interface to mock the implementation
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("bcrypt.GenerateFromPassword: %w", err), functionName)
	}

	user := &entity.User{}
	user.Email = payload.Email
	user.Fullname = payload.Fullname
	user.CryptedPassword = string(cryptedPassword)
	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateUser: %w", err), functionName)
	}

	return user, nil
}
