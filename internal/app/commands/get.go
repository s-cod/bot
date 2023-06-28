package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	dat := "DTO"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, dat)
	c.bot.Send(msg)

}

func init() {
	commands["get"] = (*Commander).Get
}
