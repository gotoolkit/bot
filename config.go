package bot

func DebugMode() OptionFunc {
	return func(bot *Bot) error {
		bot.tb.Debug = true
		return nil
	}
}
