package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Starthandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

func GetMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	chatid := update.Message.Chat.ID
	text := update.Message.Text

	switch text {
	case "/GetClothes":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Вот список вещей: ",
		})
	case "/WashedClothes":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Нажмите кнопку с цветом выбраных вещей или нажмите все вещи, чтобы выбрать самрому: ",
		})
	case "/AddClothes":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Введите список вещей которые хотите добавить в виде вещь-цвет-количество через запятую",
		})
	case "/menu":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
		})
	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Такой комманды нет",
		})
	}
}
