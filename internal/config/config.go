package config

import (
	"fmt"
	"os"
)

type Config struct {
	NanoBananaAPIKey string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("NANOBANANA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("NANOBANANA_API_KEY environment variable is not set")
	}

	return &Config{
		NanoBananaAPIKey: apiKey,
	}, nil
}
