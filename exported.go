package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Chat struct {
	tgbotapi.Chat
}
type Messager struct {
	tgbotapi.MessageConfig
}

func NewMessage(chatID int64, msg string) Messager {
	messager := tgbotapi.NewMessage(chatID, msg)
	return Messager{messager}
}
