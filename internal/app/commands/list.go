package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	dat := "Here all products:\n\n"
	for _, i := range c.productService.List() {
		dat += i.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, dat)
	c.bot.Send(msg)

}

func init() {
	commands["list"] = (*Commander).List
}
