package builders

import (
	"github.com/makehlv/tgbot/builders/callback_data"
	"github.com/makehlv/tgbot/builders/inline_keyboard"
	"github.com/makehlv/tgbot/builders/pagination"
	"github.com/makehlv/tgbot/config"
)

type Builders struct {
	CallbackData   *callbackdata.Builder
	InlineKeyboard *inlinekeyboard.Builder
	Pagination     *pagination.Builder
}

func New(cfg *config.Config) *Builders {
	return &Builders{
		CallbackData:   callbackdata.New(),
		InlineKeyboard: inlinekeyboard.New(),
		Pagination:     pagination.New(cfg, callbackdata.New(), inlinekeyboard.New()),
	}
}
