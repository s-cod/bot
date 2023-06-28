package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Чем тебе помочь: \ncommand /list",
	)
	c.bot.Send(msg)
}

func init() {
	commands["help"] = (*Commander).Help
}
