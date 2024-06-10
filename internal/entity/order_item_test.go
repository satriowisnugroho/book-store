package entity_test

import (
	"testing"

	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestOrderItemPayloadValidate(t *testing.T) {
	testcases := []struct {
		name    string
		payload *entity.OrderItemPayload
		wantErr bool
	}{
		{
			name:    "invalid quantity",
			payload: &entity.OrderItemPayload{},
			wantErr: true,
		},
		{
			name:    "success",
			payload: &entity.OrderItemPayload{Quantity: 5},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.wantErr, tc.payload.Validate() != nil)
	}
}
