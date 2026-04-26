package telegram

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	fsm "github.com/whynot00/go-telegram-fsm"
)

type Repository interface {
	InsertTable(u *entity.Thing, id int64)
	UpdateTable(u *entity.Thing, id int64)
	ReadValues(id int64) []entity.Thing
	DeleteValues(u *entity.Thing, id int64)
}

type Bot struct {
	api  *bot.Bot
	repo Repository
}

const (
	stateDefault     fsm.StateFSM = fsm.StateDefault
	stateWaitMessage fsm.StateFSM = "wait_message"
	stateColor       fsm.StateFSM = "wait_color"
	stateClothes     fsm.StateFSM = "wait_clothes"
)

func NewBot(ctx context.Context, cfg *config.Config, repo Repository) (*Bot, error) {
	slog.Debug("Запуск функции NewBot")

	mybot := &Bot{repo: repo}

	machine := fsm.New(ctx,
		fsm.WithCleanupInterval(1*time.Minute),
		fsm.WithTTL(30*time.Second),
	)

	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, mybot.StartHandler),
		bot.WithMessageTextHandler("/menu", bot.MatchTypeExact, mybot.MenuHandler),
		bot.WithMessageTextHandler("/GetClothes", bot.MatchTypeExact, mybot.GetClothesHandler),
		//bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
		//bot.WithDefaultHandler(mybot.Defaulthandler),
		bot.WithMiddlewares(fsm.Middleware(machine)),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		slog.Error("Ошибка создания бота", "error", err)
		os.Exit(1)
	}

	mybot.api = b

	mybot.Handlers()

	slog.Debug("завершение функции NewBot")
	return mybot, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.api.Start(ctx)
}

func (b *Bot) MakeButtons(chatid int64, color string) ([][]models.InlineKeyboardButton, string) {
	values := b.repo.ReadValues(chatid)
	things := entity.ThingsFromColors(values, color)

	NumOfThings := len(things)

	var numofrows int
	if NumOfThings < 4 {
		numofrows = NumOfThings + 1
	} else if NumOfThings%3 == 0 {
		numofrows = NumOfThings/3 + 1
	} else {
		numofrows = NumOfThings/3 + 2
	}
	arr := make([][]models.InlineKeyboardButton, numofrows)

	if NumOfThings > 3 {
		var j = 0
		for _, t := range things {
			arr[j] = append(arr[j], models.InlineKeyboardButton{Text: t + "-" + color, CallbackData: "buttom" + t})
			if len(arr[j]) == 3 {
				j += 1
			}
		}
	} else {
		for i, t := range things {
			arr[i] = append(arr[i], models.InlineKeyboardButton{Text: t + "-" + color, CallbackData: "buttom" + t})
		}
	}

	if len(arr) == 0 {
		return nil, "Вещей данной категории не найдено"
	} else {
		arr[len(arr)-1] = append(arr[len(arr)-1], models.InlineKeyboardButton{Text: "Выбрать", CallbackData: "buttom_selection"})
		return arr, ""
	}
}
