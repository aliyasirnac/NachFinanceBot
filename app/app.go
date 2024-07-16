package app

import (
	"context"

	"github.com/aliyasirnac/gelirgiderbot/internal/app/bot"
	"github.com/aliyasirnac/gelirgiderbot/internal/config"
	"github.com/aliyasirnac/gelirgiderbot/internal/db"

	"github.com/aliyasirnac/gelirgiderbot/internal/loggerx"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg       *config.Config
	logger    logrus.FieldLogger
	cancel    context.CancelFunc
	botStatus bool
}

func New(cfg *config.Config) *App {
	logger := loggerx.New(cfg.App.Log)
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.logger.Info("starting app")

	a.logger.Info("database starting")
	database, err := db.New(*a.cfg)

	if err != nil {
		return err
	}

	err = database.AutoMigrate(&db.User{})
	if err != nil {
		return err
	}

	a.logger.Info("bot starting")
	b := bot.New(a.cfg.App.Token)

	defer func() {
		if err = b.Run(ctx); err != nil {
			loggerx.ExitOnError(err, "bot çalıştırılırken bir hata oluştu")
		}
	}()

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("stopping app")

	return nil
}
