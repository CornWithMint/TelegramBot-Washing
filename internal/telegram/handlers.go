package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Starthandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddNewClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

func GetCloteshandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddNewClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}
