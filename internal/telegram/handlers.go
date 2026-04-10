package telegram

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	fsm "github.com/whynot00/go-telegram-fsm"
)

func (b *Bot) Handlers() {
	slog.Debug("Запуск функции Handlers")
	// ADDCLOTHES HANDLER 1
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "/AddClothes", bot.MatchTypeExact,
		func(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
			chatid := update.Message.Chat.ID

			f := fsm.FromContext(ctx)
			f.Transition(ctx, stateWaitMessage)
			BotApi.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatid,
				Text:   "Введите список вещей которые хотите добавить в виде вещь-цвет-количество через запятую",
			})
		},
		fsm.WithStates(stateDefault),
	)

	// WashedClothesHandler 1
	// Сделать возможность мультивыбора
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "/WashedClothes", bot.MatchTypeExact,
		func(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
			chatid := update.Message.Chat.ID

			f := fsm.FromContext(ctx)
			f.Transition(ctx, stateColor)

			kb := models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{Text: "Черные", CallbackData: "button_1"},
						{Text: "Белые", CallbackData: "button_2"},
					},
					{
						{Text: "Цветные", CallbackData: "button_3"},
						{Text: "Все", CallbackData: "button_4"},
					},
				},
			}

			BotApi.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      chatid,
				Text:        "Выберите каких цветов постирали вещи или все для вывода всех вещей",
				ReplyMarkup: kb,
			})
		},
		fsm.WithStates(stateDefault),
	)

	b.api.RegisterHandler(bot.HandlerTypeCallbackQueryData, "button", bot.MatchTypePrefix,
		func(ctx context.Context, BotApi *bot.Bot, update *models.Update) {

			chatid := update.Message.Chat.ID
			data := update.CallbackQuery.Data

			f := fsm.FromContext(ctx)
			f.Transition(ctx, stateClothes)

			switch data {
			case "button_1":
				b.api.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Черный мальчик",
				})
			case "button_2":
				b.api.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Белый мальчик",
				})
			case "button_3":
				b.api.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Цветной мальчик",
				})
			case "button_4":
				b.api.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Все мальчики",
				})
			}

			BotApi.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: update.CallbackQuery.ID,
				// Text: "Обрабатываю...", // Необязательное всплывающее уведомление
			})
		},
		fsm.WithStates(stateColor),
	)

	//ADDCLOTHES HANDLER 2
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix,
		func(ctx context.Context, BotApi *bot.Bot, update *models.Update) {

			NewText := update.Message.Text
			chatid := update.Message.Chat.ID

			f := fsm.FromContext(ctx)
			f.Finish(ctx)
			clothes, err := entity.StringToUserArr(NewText, chatid)
			if err != nil {
				BotApi.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Данные введены не верно",
				})
			} else {
				for _, clothe := range clothes {
					b.repo.UpdateTable(&clothe, chatid)
				}
				BotApi.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatid,
					Text:   "Вещи загружены:",
				})
			}

		},
		fsm.WithStates(stateWaitMessage),
	)
	slog.Debug("Завершение функции Handlers")
}

func (b *Bot) StartHandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
	BotApi.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

func (b *Bot) GetClothesHandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {

	chatid := update.Message.Chat.ID
	NewText := entity.UsersArrToString(b.repo.ReadValues(chatid))

	BotApi.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   fmt.Sprintf("Вот список вещей: \n %s", NewText),
	})
	b.repo.ReadValues(chatid)
}

func (b *Bot) MenuHandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
	chatid := update.Message.Chat.ID
	BotApi.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

func (b *Bot) DefaultHandler(ctx context.Context, BotApi *bot.Bot, update *models.Update) {
	chatid := update.Message.Chat.ID
	BotApi.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   "Такой комманды нет"})
}
