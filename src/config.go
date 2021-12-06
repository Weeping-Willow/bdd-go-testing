package src

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiToken string
}

func NewConfig(path ...string) (*Config, error) {
	if err := godotenv.Load(path...); err != nil {
		return nil, err
	}
	conf := &Config{
		ApiToken: getEnvWithDefault("API_TOKEN", ""),
	}

	if conf.ApiToken == "" {
		return nil, errors.New("no API_TOKEN found")
	}
	return conf, nil
}

func getEnvWithDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
