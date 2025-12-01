package pagination

import (
	callbackdata "github.com/makehlv/tgbot/builders/callback_data"
	inlinekeyboard "github.com/makehlv/tgbot/builders/inline_keyboard"
	"github.com/makehlv/tgbot/config"
)

type Builder struct {
	callbackDataBuilder *callbackdata.Builder
	keyboardBuilder     *inlinekeyboard.Builder
	BaseOffset          int
	Limit               int
}

func New(cfg *config.Config, callbackDataBuilder *callbackdata.Builder, keyboardBuilder *inlinekeyboard.Builder) *Builder {
	return &Builder{
		callbackDataBuilder: callbackDataBuilder,
		keyboardBuilder:     keyboardBuilder,
		BaseOffset:          5,
		Limit:               cfg.TelegramListLimitLen,
	}
}
