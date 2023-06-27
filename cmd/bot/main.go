package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/s-cod/bot/internal/service/product"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	productService := product.NewService()
	commander := NewCommander(bot, productService)

	for update := range updates {

		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "help":
			commander.helpCommand(update.Message)
		case "list":
			commander.listCommand(update.Message)
		default:
			commander.defaultHandler(update.Message)
		}

	}
}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c Commander) helpCommand(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Чем тебе помочь")
	c.bot.Send(msg)
}
func (c Commander) listCommand(inputMessage *tgbotapi.Message) {
	dat := "Here all products:\n\n"
	for _, i := range c.productService.List() {
		dat += i.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, dat)
	c.bot.Send(msg)
}

func (c Commander) defaultHandler(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ты написал: "+inputMessage.Text)
	c.bot.Send(msg)
}
