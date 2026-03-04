package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
)

func StartBot(token string, ctx context.Context) {

	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, Starthandler),
		bot.WithDefaultHandler(GetMessageHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота: ", err)
	}
	fmt.Println("Бот cоздан")
	b.Start(ctx)
	fmt.Println("Бот запущен")

}
