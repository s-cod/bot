package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	dat := "Here all products:\n\n"
	for _, i := range c.productService.List() {
		dat += i.Title + "\n"
	}

	data, err := json.Marshal(CommandData{
		Offset: 10,
	})

	if err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, dat)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(data)),
		),
	)
	c.bot.Send(msg)
}
