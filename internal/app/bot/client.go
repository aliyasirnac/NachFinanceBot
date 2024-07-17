package bot

import (
	"context"
	"github.com/aliyasirnac/gelirgiderbot/internal/db"
	"time"

	"github.com/aliyasirnac/gelirgiderbot/internal/loggerx"
	"gopkg.in/telebot.v3"
)

type Client struct {
	appToken string
	bot      *telebot.Bot
	db       db.Service
}

func New(appToken string, db db.Service) *Client {
	return &Client{
		appToken: appToken,
		db:       db,
	}
}

func (c *Client) Run(ctx context.Context) error {
	pref := telebot.Settings{
		Token:  c.appToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		loggerx.ExitOnError(err, "Bot oluşturulurken hata oluştu")
		return err
	}

	c.bot = b
	Register(b, c.db)

	go func() {
		<-ctx.Done()
		b.Stop()
	}()

	b.Start()

	return nil
}

func (c *Client) Stop() {
	if c.bot != nil {
		c.bot.Stop()
	}
}
