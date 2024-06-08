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
	"golang.org/x/crypto/bcrypt"
)

// UserUsecaseInterface define contract for user related functions to usecase
type UserUsecaseInterface interface {
	CreateUser(ctx context.Context, payload *entity.RegisterPayload) (*entity.User, error)
	Login(ctx context.Context, payload *entity.LoginPayload) (*entity.LoginResponse, error)
}

type UserUsecase struct {
	userRepo repo.UserRepositoryInterface
}

func NewUserUsecase(ur repo.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
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
		return nil, errors.Wrap(fmt.Errorf("uc.userRepo.GetUserByEmail: %w", err), functionName)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.CryptedPassword), []byte(payload.Password))
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
	tokenStr, _ := token.SignedString([]byte("secret"))

	resp := &entity.LoginResponse{
		AccessToken: tokenStr,
	}

	return resp, nil
}
