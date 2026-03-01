package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/telegram"

	"github.com/joho/godotenv"
)

func main() {
	// Контекст для graseful shotdown бота после interrupt
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment")
	}

	// Загружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//Запускаем бота
	telegram.StartBot(cfg.BotToken, ctx)
}
