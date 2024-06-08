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

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name     string
		ctx      context.Context
		payload  *entity.UserPayload
		rUserErr error
		wantErr  bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:    "invalid payload",
			ctx:     context.Background(),
			payload: &entity.UserPayload{Email: "foo"},
			wantErr: true,
		},
		{
			name:     "failed to create user",
			ctx:      context.Background(),
			payload:  &entity.UserPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			rUserErr: errors.New("error create user"),
			wantErr:  true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			payload: &entity.UserPayload{Email: "foo@bar.com", Fullname: "Foo Bar", Password: "12345"},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := &testmock.UserRepositoryInterface{}
			userRepo.On("CreateUser", mock.Anything, mock.Anything).Return(tc.rUserErr)

			uc := usecase.NewUserUsecase(userRepo)
			_, err := uc.CreateUser(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
