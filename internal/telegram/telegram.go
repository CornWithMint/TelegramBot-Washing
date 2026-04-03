package telegram

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/go-telegram/bot"
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
	stateDefault     fsm.StateFSM = "default"
	stateWaitMessage fsm.StateFSM = "wait_message"
	stateColor       fsm.StateFSM = "wait_color"
)

func NewBot(ctx context.Context, cfg *config.Config, repo Repository) (*Bot, error) {

	mybot := &Bot{repo: repo}

	machine := fsm.New(ctx,
		fsm.WithCleanupInterval(1*time.Minute),
		fsm.WithTTL(30*time.Second),
	)

	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, mybot.StartHandler),
		bot.WithMessageTextHandler("/menu", bot.MatchTypeExact, mybot.MenuHandler),
		bot.WithMessageTextHandler("/GetClothes", bot.MatchTypeExact, mybot.GetClothesHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
		//bot.WithDefaultHandler(mybot.Defaulthandler),
		bot.WithMiddlewares(fsm.Middleware(machine)),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота: ", err)
	}

	mybot.api = b

	mybot.Handlers()

	return mybot, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.api.Start(ctx)
	fmt.Println("Бот запущен")
}
