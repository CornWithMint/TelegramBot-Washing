package config

import (
	"os"
)

type Config struct {
	BotToken string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}

	return cfg, nil
}
