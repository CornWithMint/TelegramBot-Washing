package telegram

import (
	"context"
	"fmt"
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
	UpdateTable(u *entity.User, id int64)
	ReadValues(id int64) []entity.User
	DeleteValues()
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

func (b *Bot) MakeButtons(chatid int64, color string) [][]models.InlineKeyboardButton {
	arr := make([][]models.InlineKeyboardButton, 0)

	values := b.repo.ReadValues(chatid)
	things := entity.ThingsFromColors(values, color)
	NumOfThings := len(things)

	if NumOfThings < 6 {
		for _, t := range things {
			arr = append(arr, []models.InlineKeyboardButton{{Text: t, CallbackData: "button_black"}})
		}
	}

	fmt.Println(arr)
	return arr
}
