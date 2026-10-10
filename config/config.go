package config

import (
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	URL string
}
type Config struct {
	Database database
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
}
