package entity_test

import (
	"testing"

	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testcases := []struct {
		name    string
		payload *entity.RegisterPayload
		wantErr bool
	}{
		{
			name: "invalid email",
			payload: &entity.RegisterPayload{
				Email: "foo",
			},
			wantErr: true,
		},
		{
			name: "invalid fullname",
			payload: &entity.RegisterPayload{
				Email:    "foo@bar.com",
				Fullname: "",
			},
			wantErr: true,
		},
		{
			name: "invalid password length",
			payload: &entity.RegisterPayload{
				Email:    "foo@bar.com",
				Fullname: "Foo Bar",
				Password: "123",
			},
			wantErr: true,
		},
		{
			name: "success",
			payload: &entity.RegisterPayload{
				Email:    "foo@bar.com",
				Fullname: "Foo Bar",
				Password: "12345",
			},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.wantErr, tc.payload.Validate() != nil)
	}
}
