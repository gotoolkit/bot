package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	tb *tgbotapi.BotAPI
}

type OptionFunc func(*Bot) error

func NewTelegramBot(botAuthToken string, opts ...OptionFunc) (*Bot, error) {

	tb, err := tgbotapi.NewBotAPI(botAuthToken)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		tb: tb,
	}

	for _, opt := range opts {
		err := opt(bot)
		if err != nil {
			return nil, err
		}
	}
	return bot, nil
}

func (b *Bot) GetUpdates(timeout int) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(1000)
	u.Timeout = timeout

	updates, err := b.tb.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return updates, nil
}

func (b *Bot) Send(messager Messager) (tgbotapi.Message, error) {
	return b.tb.Send(messager)
}
