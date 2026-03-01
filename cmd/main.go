package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/CornWithMint/TelegramBot-Washing/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

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

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота")
	}
	b.Start(ctx)
}
