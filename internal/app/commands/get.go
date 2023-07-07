package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	// dat := "DTO"

	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("нет данных")
		return
	}

	getData, err := c.productService.Get(arg)
	if err != nil {
		log.Println("Нет данных")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("command: %s, arg: %v", "get", getData.Title),
	)
	c.bot.Send(msg)

}
