package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     uint16 `env:"PORT,default=9999"`
	Env      string `env:"ENV"`
	LogLevel string `env:"LOG_LEVEL,default=debug"`
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
