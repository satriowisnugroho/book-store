package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/test/fixture"
	testmock "github.com/satriowisnugroho/book-store/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		payload   *entity.RegisterPayload
		hasherRes string
		hasherErr error
		rUserErr  error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:    "invalid payload",
			ctx:     context.Background(),
			payload: &entity.RegisterPayload{Email: "foo"},
			wantErr: true,
		},
		{
			name:      "failed to hashing the password",
			ctx:       context.Background(),
			payload:   &entity.RegisterPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			hasherErr: errors.New("error hashing password"),
			wantErr:   true,
		},
		{
			name:     "failed when email is duplicate",
			ctx:      context.Background(),
			payload:  &entity.RegisterPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			rUserErr: response.ErrInvalidEmail,
			wantErr:  true,
		},
		{
			name:     "failed to create user",
			ctx:      context.Background(),
			payload:  &entity.RegisterPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			rUserErr: errors.New("error create user"),
			wantErr:  true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			payload: &entity.RegisterPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			passwordHasher := &testmock.PasswordHasher{}
			passwordHasher.On("GenerateFromPassword", mock.Anything).Return(tc.hasherRes, tc.hasherErr)

			userRepo := &testmock.UserRepositoryInterface{}
			userRepo.On("CreateUser", mock.Anything, mock.Anything).Return(tc.rUserErr)

			uc := usecase.NewUserUsecase("", passwordHasher, userRepo)
			_, err := uc.CreateUser(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	testcases := []struct {
		name      string
		ctx       context.Context
		rUserRes  *entity.User
		rUserErr  error
		hasherErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:     "failed to get user by email",
			ctx:      context.Background(),
			rUserErr: errors.New("error get user by email"),
			wantErr:  true,
		},
		{
			name: "failed to compare password",
			ctx:  context.Background(),
			rUserRes: &entity.User{
				ID:              123,
				Email:           "foo@bar.com",
				Fullname:        "Foo Bar",
				CryptedPassword: string(hashedPassword),
			},
			hasherErr: errors.New("error compare password"),
			wantErr:   true,
		},
		{
			name: "success",
			ctx:  context.Background(),
			rUserRes: &entity.User{
				ID:              123,
				Email:           "foo@bar.com",
				Fullname:        "Foo Bar",
				CryptedPassword: string(hashedPassword),
			},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			passwordHasher := &testmock.PasswordHasher{}
			passwordHasher.On("CompareHashAndPassword", mock.Anything, mock.Anything).Return(tc.hasherErr)

			userRepo := &testmock.UserRepositoryInterface{}
			userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(tc.rUserRes, tc.rUserErr)

			uc := usecase.NewUserUsecase("", passwordHasher, userRepo)
			_, err := uc.Login(tc.ctx, &entity.LoginPayload{Password: "password123"})
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
