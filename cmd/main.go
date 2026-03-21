package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/database"
	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/CornWithMint/TelegramBot-Washing/internal/telegram"

	"github.com/joho/godotenv"
)

func main() {
	// Контекст для graseful shotdown бота после interrupt
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	//Загрузка env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment")
	}

	// Загружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//Открываем файл для логгов
	file, err := os.OpenFile(cfg.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Не удалось загрузить файл для логирования", err)
	}
	log.SetOutput(file)

	db, err := database.NewSqliteRepo(cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dad")
	//Запускаем бота
	bot, err := telegram.NewBot(cfg, db)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	bot.Start(ctx)

	Users := &[]entity.User{
		{Id: 0, Thing: "Jeans", Color: "Black", Number: 1},
	}
	for _, user := range *Users {
		db.UpdateTable(&user)
	}

}
