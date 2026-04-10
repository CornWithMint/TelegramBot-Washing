package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/database"
	"github.com/CornWithMint/TelegramBot-Washing/internal/telegram"

	"github.com/joho/godotenv"
)

func main() {
	//Создаем логи
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	// Контекст для graseful shotdown бота после interrupt
	slog.Info("Создание контекста")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	//Загрузка env файла
	slog.Info("Запуск env файла")

	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Ошибка открытия .env файла", "error", err)
	}

	// Загружаем конфиг
	slog.Info("Загрузка конфига")
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Ошибка загрузки конфига", "error", err)
		os.Exit(1)
	}

	//Запуск БД
	slog.Info("Загрузка БД")
	db := database.NewSqliteRepo(cfg)

	//Запускаем бота
	slog.Info("Запуск бота")
	bot, err := telegram.NewBot(ctx, cfg, db)
	if err != nil {
		slog.Error("Ошибка запуска бота", "error", err)
		os.Exit(1)
	}
	slog.Info("Бот запущен")
	bot.Start(ctx)
}
