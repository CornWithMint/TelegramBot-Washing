package telegram

import (
	"context"

	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *Bot) Starthandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
	BotApi.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

// func WashedClothesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
// 	inlinebuton := inline.New(b).
// 		Row().
// 		Button("Белые", []byte("1-1")).
// 		Row().
// 		Button("Черные", []byte("2-1")).
// 		Row().
// 		Button("Цветные", []byte("3-1")).
// 		Row().
// 		Button("Все", []byte("4-1"))

// 	b.SendMessage(ctx, &bot.SendMessageParams{
// 		ChatID:      update.Message.Chat.ID,
// 		Text:        "Нажмите кнопку с цветом выбраных вещей или нажмите все вещи, чтобы выбрать самрому: ",
// 		ReplyMarkup: inlinebuton,
// 	})
// }

func (b *Bot) GetMessageHandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	chatid := update.Message.Chat.ID
	text := update.Message.Text

	switch text {
	case "/GetClothes":
		NewText := entity.UsersArrToString(b.repo.ReadValues(0))

		BotApi.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Вот список вещей: ",
		})
		BotApi.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   NewText,
		})
		b.repo.ReadValues(chatid)
	case "/AddClothes":
		BotApi.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Введите список вещей которые хотите добавить в виде вещь-цвет-количество через запятую",
		})
		if update.Message.Text != "" {

		}

	case "/menu":
		BotApi.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
		})
	default:
		BotApi.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Такой комманды нет",
		})
	}
}
