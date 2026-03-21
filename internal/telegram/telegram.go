package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	"github.com/go-telegram/bot"
)

type Repository interface {
	UpdateTable(u *entity.User)
	ReadValues(id int)
	DeleteValues()
}

type Bot struct {
	api  *bot.Bot
	repo Repository
}

func NewBot(cfg *config.Config, repo Repository) (*Bot, error) {

	mybot := &Bot{repo: repo}
	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, mybot.Starthandler),
		//bot.WithMessageTextHandler("/WashedClothes", bot.MatchTypeExact, WashedClothesHandler),
		bot.WithDefaultHandler(mybot.GetMessageHandler),
	}
	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		log.Fatal("Ошибка создания бота: ", err)
	}
	mybot.api = b

	return mybot, nil
}

func Format(clothes string, id int) ([]entity.User, error) {
	var res [][]string
	var things []entity.User

	for _, val := range strings.Split(clothes, ",") {
		if val == "" || !strings.Contains(val, "-") {
			return nil, errors.New("Значени введены не верно")
		} else {
			strings.Trim(val, " ")
			res = append(res, strings.Split(val, "-"))
		}
	}

	for i := range res {
		num, _ := strconv.Atoi(res[i][2])

		u := &entity.User{
			Id:     id,
			Thing:  res[i][0],
			Color:  res[i][1],
			Number: num,
		}
		things = append(things, *u)

	}

	return things, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.api.Start(ctx)
	fmt.Println("Бот запущен")
}
