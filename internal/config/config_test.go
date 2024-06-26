package config_test

import (
	"os"
	"testing"

	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert.Panics(t, func() { config.NewConfig() }, "the config panics")

	setupEnv()
	cfg := config.NewConfig()

	assert.NotEmpty(t, cfg)
}

func setupEnv() {
	os.Setenv("DATABASE_USERNAME", "postgres")
	os.Setenv("DATABASE_PASSWORD", "postgres")
	os.Setenv("DATABASE_NAME", "book_store_development")
}
