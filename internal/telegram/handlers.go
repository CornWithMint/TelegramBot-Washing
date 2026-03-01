package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
