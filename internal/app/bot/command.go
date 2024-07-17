package bot

import (
	"fmt"
	"github.com/aliyasirnac/gelirgiderbot/internal/db"
	"gopkg.in/telebot.v3"
	"log"
)

type CommandFunc func(ctx telebot.Context, service db.Service) error

func helloCommand() CommandFunc {
	return func(ctx telebot.Context, service db.Service) error {
		user := ctx.Sender()
		return ctx.Send(fmt.Sprintf("ID: %v Username: %s", user.ID, user.Username))
	}
}

func helpCommand() CommandFunc {
	return func(ctx telebot.Context, _ db.Service) error {

		return ctx.Send("Help")
	}
}

func startCommand() CommandFunc {
	return func(ctx telebot.Context, service db.Service) error {
		msg := "Welcome to NachFinanceBot! This bot helps you track your income and expenses. Use /help to see available commands."

		u := ctx.Sender()
		_, err := service.GetUserByBotId(u.ID)

		// user does not exist
		if err != nil {
			user := db.User{
				BotId:     u.ID,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Goals:     nil,
			}
			err = service.AddUser(user)
			if err != nil {
				log.Println("failed to add user:", err)

				return err
			}
			return ctx.Send(msg)
		}
		log.Println("user already exists:", err)

		return ctx.Send(msg)
	}
}

func Register(b *telebot.Bot, service db.Service) {
	b.Handle("/hello", func(ctx telebot.Context) error {
		return helloCommand()(ctx, service)
	})
	b.Handle("/help", func(ctx telebot.Context) error {
		return helpCommand()(ctx, service)
	})
	b.Handle("/start", func(ctx telebot.Context) error {
		return startCommand()(ctx, service)
	})
}
