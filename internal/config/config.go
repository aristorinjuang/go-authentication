package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Database struct {
	Source          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type TokenDuration struct {
	Access time.Duration
}

type TokenSecret struct {
	Access  string
	Refresh string
}

type Config struct {
	Database      Database
	Origins       string
	Port          string
	TokenDuration TokenDuration
	TokenSecret   TokenSecret
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			Source:          os.Getenv("DATABASE_SOURCE"),
			MaxIdleConns:    cast.ToInt(os.Getenv("DATABASE_MAX_IDLE_CONNS")),
			MaxOpenConns:    cast.ToInt(os.Getenv("DATABASE_MAX_OPEN_CONNS")),
			ConnMaxLifetime: time.Duration(cast.ToInt(os.Getenv("DATABASE_CONN_MAX_LIFETIME_MINUTES"))) * time.Minute,
		},
		Origins: os.Getenv("ORIGINS"),
		Port:    os.Getenv("PORT"),
		TokenDuration: TokenDuration{
			Access: time.Duration(cast.ToInt(os.Getenv("ACCESS_TOKEN_EXPIRES_IN_HOURS"))) * time.Hour,
		},
		TokenSecret: TokenSecret{
			Access:  os.Getenv("ACCESS_TOKEN_SECRET"),
			Refresh: os.Getenv("REFRESH_TOKEN_SECRET"),
		},
	}, nil
}
