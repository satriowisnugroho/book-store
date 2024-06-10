package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
	repo "github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/pkg/auth"
)

// UserUsecaseInterface define contract for user related functions to usecase
type UserUsecaseInterface interface {
	CreateUser(ctx context.Context, payload *entity.RegisterPayload) (*entity.User, error)
	Login(ctx context.Context, payload *entity.LoginPayload) (*entity.LoginResponse, error)
}

type UserUsecase struct {
	jwtSecret      string
	passwordHasher auth.PasswordHasher
	userRepo       repo.UserRepositoryInterface
}

func NewUserUsecase(js string, ph auth.PasswordHasher, ur repo.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		jwtSecret:      js,
		passwordHasher: ph,
		userRepo:       ur,
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, payload *entity.RegisterPayload) (*entity.User, error) {
	functionName := "UserUsecase.CreateUser"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	cryptedPassword, err := uc.passwordHasher.GenerateFromPassword(payload.Password)
	if err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.passwordHasher.GenerateFromPassword: %w", err), functionName)
	}

	user := &entity.User{}
	user.Email = payload.Email
	user.Fullname = payload.Fullname
	user.CryptedPassword = cryptedPassword
	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		if _, ok := err.(response.CustomError); ok {
			return nil, err
		}

		return nil, errors.Wrap(fmt.Errorf("uc.userRepo.CreateUser: %w", err), functionName)
	}

	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, payload *entity.LoginPayload) (*entity.LoginResponse, error) {
	functionName := "UserUsecase.Login"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	user, err := uc.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		if _, ok := err.(response.CustomError); ok {
			return nil, err
		}

		return nil, errors.Wrap(fmt.Errorf("uc.userRepo.GetUserByEmail: %w", err), functionName)
	}

	err = uc.passwordHasher.CompareHashAndPassword(user.CryptedPassword, payload.Password)
	if err != nil {
		return nil, response.ErrInvalidPassword
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"fullname": user.Fullname,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(uc.jwtSecret))

	resp := &entity.LoginResponse{
		AccessToken: tokenStr,
	}

	return resp, nil
}
