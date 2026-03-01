package telegram

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
)

func StartBot(token string, ctx context.Context) {

	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, Starthandler),
		bot.WithMessageTextHandler("/menu", bot.MatchTypeExact, Starthandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота")
	}

	b.Start(ctx)
}
