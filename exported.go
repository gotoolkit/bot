package bot

func DebugMode(debug bool) OptionFunc {
	return func(bot *Bot) error {
		bot.tg.Debug = debug
		return nil
	}
}

func ParseMode(parseMode string) OptionFunc {
	return func(bot *Bot) error {
		bot.parseMode = parseMode
		return nil
	}
}

func WithChatIDs(chatIDs ...int64) OptionFunc {
	return func(bot *Bot) error {
		bot.chatIDs = append(bot.chatIDs, chatIDs...)
		return nil
	}
}
