package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseDSN string
}

func Load() (*Config, error) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return nil, fmt.Errorf("SERVER_PORT is required")
	}

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_DSN is required")
	}

	return &Config{
		ServerPort:  port,
		DatabaseDSN: dsn,
	}, nil
}
