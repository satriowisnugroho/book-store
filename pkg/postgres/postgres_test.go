package postgres_test

import (
	"testing"

	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgres(t *testing.T) {
	testcases := []struct {
		name    string
		config  config.DatabaseConfig
		wantErr bool
	}{
		{
			name:    "correct config",
			config:  config.DatabaseConfig{Driver: "postgres"},
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  config.DatabaseConfig{Driver: "invalid"},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		postgresDb, err := postgres.NewPostgres(&tc.config)

		assert.NotNil(t, postgresDb)
		if tc.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
