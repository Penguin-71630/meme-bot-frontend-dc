package config

import (
	"errors"
	"os"
)

type Config struct {
	DiscordToken string
	APIBaseURL   string
	BotPrefix    string
}

func Load() (*Config, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("DISCORD_BOT_TOKEN is required")
	}

	apiURL := os.Getenv("API_BASE_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080" // Default
	}

	prefix := os.Getenv("BOT_PREFIX")
	if prefix == "" {
		prefix = "!" // Default prefix
	}

	return &Config{
		DiscordToken: token,
		APIBaseURL:   apiURL,
		BotPrefix:    prefix,
	}, nil
}
