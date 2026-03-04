package config

import (
	"os"
)

type Config struct {
	BotToken string
	BdPath   string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		BdPath:   os.Getenv("BD_PATH"),
	}

	return cfg, nil
}
