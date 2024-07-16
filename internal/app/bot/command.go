package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

type BotCommandFunc func(ctx telebot.Context) error

func helloCommand() BotCommandFunc {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		return ctx.Send(fmt.Sprintf("ID: %v Username: %s", user.ID, user.Username))
	}
}

func helpCommand() BotCommandFunc {
	return func(ctx telebot.Context) error {
		return ctx.Send("Help")
	}
}

func StartCommand() BotCommandFunc {
	return func(ctx telebot.Context) error {
		msg := "Welcome to NachFinanceBot! This bot helps you track your income and expenses. Use /help to see available commands."
		return ctx.Send(msg)
	}
}

func Register(b *telebot.Bot) {
	b.Handle("/hello", telebot.HandlerFunc(helloCommand()))
	b.Handle("/help", telebot.HandlerFunc(helpCommand()))
	b.Handle("/start", telebot.HandlerFunc(StartCommand()))
}
