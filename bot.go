package telegram

import (
	"log/slog"

	"github.com/makehlv/tgbot/builders"
	"github.com/makehlv/tgbot/client"
	"github.com/makehlv/tgbot/config"
	"github.com/makehlv/tgbot/fsm"
	"github.com/makehlv/tgbot/server"
)

type Bot struct {
	logger   *slog.Logger
	cfg      *config.Config
	fsm      *fsm.FSM
	client   *client.Client
	builders *builders.Builders
	server   *server.Server
}

func New(cfg *config.Config, logger *slog.Logger, stateStore fsm.StateStore) *Bot {
	updateHandler := fsm.New(logger, stateStore)
	telegramClient := client.New(cfg, logger)
	builders := builders.New(cfg)
	scoper := NewScoper(updateHandler, telegramClient, builders)
	server := server.New(cfg, logger, scoper)
	return &Bot{
		logger:   logger,
		cfg:      cfg,
		fsm:      updateHandler,
		client:   telegramClient,
		builders: builders,
		server:   server,
	}
}
