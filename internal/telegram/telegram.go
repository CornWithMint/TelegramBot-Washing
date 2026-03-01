package telegram

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
)

func StartBot(token string, ctx context.Context) {

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота")
	}

	b.Start(ctx)
}
