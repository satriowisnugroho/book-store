package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Port           uint16 `env:"PORT,default=9999"`
	Env            string `env:"ENV"`
	LogLevel       string `env:"LOG_LEVEL,default=debug"`
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string `env:"DATABASE_DRIVER,default=postgres"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Name     string `env:"DATABASE_NAME,required"`
	SslMode  string `env:"DATABASE_SSL_MODE,default=disable"`
	Pool     int    `env:"DATABASE_POOL,default=50"`
}

func NewConfig() *Config {
	var cfg Config
	godotenv.Load(".env")
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
