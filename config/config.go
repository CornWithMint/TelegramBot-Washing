package config

import (
	"os"
)

type Config struct {
	BotToken string
	BdPath   string
	LogPath  string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		BdPath:   os.Getenv("BD_PATH"),
		LogPath:  os.Getenv("LOG_PATH"),
	}

	return cfg, nil
}
