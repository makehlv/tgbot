package telegram

import (
	"github.com/makehlv/tgbot/builders"
	"github.com/makehlv/tgbot/client"
	"github.com/makehlv/tgbot/models"
)

type Scope struct {
	client   *client.Client
	builders *builders.Builders
	update   *models.Update
}

func NewScope(client *client.Client, builders *builders.Builders, update *models.Update) *Scope {
	return &Scope{
		client:   client,
		builders: builders,
		update:   update,
	}
}
