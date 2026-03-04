package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/database"
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

	// Открываем базу данных
	db, err := sql.Open("sqlite3", cfg.BdPath)
	if err != nil {
		log.Fatal("Ошибка открытия бд", err)
	}
	database.CreateTable(db)
	database.ReadValues(0, db)

	defer db.Close()
	//Запускаем бота
	telegram.StartBot(cfg.BotToken, ctx)

}
