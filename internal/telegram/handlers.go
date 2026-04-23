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
		b.AddClothesHandler,
		fsm.WithStates(fsm.StateAny),
	)

	// WashedClothesHandler (Создает кнопки с цветами)
	// Сделать возможность мультивыбора
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "/WashedClothes", bot.MatchTypeExact,
		b.WashedClothesHandler,
		fsm.WithStates(fsm.StateAny),
	)

	// WashdClothesHancler 2 (обработка нажания на кнопку цвета)
	b.api.RegisterHandler(bot.HandlerTypeCallbackQueryData, "", bot.MatchTypePrefix,
		b.WashedAnswer,
		fsm.WithStates(fsm.StateAny),
	)

	//ADDCLOTHES HANDLER 2
	// !!ДОБАВИТЬ ВОЗМОЖНОСТЬ ЗАГРУЗИТЬ ВЕЩИ ЧЕРЕЗ ФАЙЛ СО СПИСКОМ ЭТИХ ВЕЩЕЙ
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix,
		b.ClothesWaitHandler,
		fsm.WithStates(stateWaitMessage),
	)
	slog.Debug("Завершение функции Handlers")
}

// Обработка /start
func (b *Bot) StartHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

// Обработка /GetClothes
func (b *Bot) GetClothesHandler(ctx context.Context, api *bot.Bot, update *models.Update) {

	chatid := update.Message.Chat.ID
	NewText := entity.UsersArrToString(b.repo.ReadValues(chatid))

	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   fmt.Sprintf("Вот список вещей: \n %s", NewText),
	})
	b.repo.ReadValues(chatid)
}

// Обработка /AddClothes
func (b *Bot) AddClothesHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	chatid := update.Message.Chat.ID

	f := fsm.FromContext(ctx)

	f.Transition(ctx, stateWaitMessage)
	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   "Введите список вещей которые хотите добавить в виде вещь-цвет-количество через запятую",
	})
}

// Ожидание Одежды после /AddClothes
func (b *Bot) ClothesWaitHandler(ctx context.Context, api *bot.Bot, update *models.Update) {

	NewText := update.Message.Text
	chatid := update.Message.Chat.ID

	f := fsm.FromContext(ctx)
	f.Finish(ctx)
	clothes, err := entity.StringToUserArr(NewText, chatid)
	if err != nil {
		api.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Данные введены не верно",
		})
	} else {
		for _, clothe := range clothes {
			b.repo.UpdateTable(&clothe, chatid)
		}
		api.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Вещи загружены:",
		})
	}

}

// Обработка /WashedClothes с кнопками цвета
func (b *Bot) WashedClothesHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
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

	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatid,
		Text:        "Выберите каких цветов постирали вещи или все для вывода всех вещей",
		ReplyMarkup: kb,
	})
}

// Ожидание кнопки с цветом после /WashedClothes
func (b *Bot) WashedAnswer(ctx context.Context, api *bot.Bot, update *models.Update) {
	slog.Debug("Функция обрабатывающая кнопки начала работать")

	callback := update.CallbackQuery
	if callback == nil {
		slog.Debug("CallBack нулевой")
		return
	}
	f := fsm.FromContext(ctx)
	f.Transition(ctx, fsm.StateDefault)

	data := callback.Data
	chatid := callback.From.ID

	switch data {
	case "button_1":
		slog.Debug("Кнопка")
		b.ColorSelectionHandler(ctx, api, chatid, "black")
	case "button_2":
		b.ColorSelectionHandler(ctx, api, chatid, "white")
	case "button_3":
		b.ColorSelectionHandler(ctx, api, chatid, "colored")
	case "button_4":
		b.ColorSelectionHandler(ctx, api, chatid, "All")
	case "button_black":
		api.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatid,
			Text:   "Срок стирки обновлен",
		})
	}

	api.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "Обрабатываю...", // Необязательное всплывающее уведомление
	})
}

func (b *Bot) ColorSelectionHandler(ctx context.Context, api *bot.Bot, chatid int64, color string) {
	slog.Debug("ColorSelectionHandler Начал работу")
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: b.MakeButtons(chatid, color),
	}

	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatid,
		Text:        "Выберите какие вещи постирали",
		ReplyMarkup: kb,
	})
}

// Обработка /menu
func (b *Bot) MenuHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	chatid := update.Message.Chat.ID
	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   "Добро пожаловать в бота который умеет считать сколько дней не стирались вещи. Введи \n /GetClothes - Чтобы посмотреть весь список вещей \n /WashedClothes - чтобы выбрать постиранные вещи \n /AddClothes - Чтобы ввести список вещей или добавить новую вещь",
	})
}

// Обработка неизвестной комманды
func (b *Bot) DefaultHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	chatid := update.Message.Chat.ID
	api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatid,
		Text:   "Такой комманды нет"})
}
