package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlad1sat/bot/internal/service/product"
	"log"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		panicValue := recover()
		if panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message == nil {
		return
	}
	msg := update.Message
	switch update.Message.Command() {
	case "help":
		c.HelpCommand(msg)
	case "list":
		c.ListCommand(msg)
	case "get":
		c.Get(msg)
	default:
		c.DefaultBehavior(msg)
	}
}
