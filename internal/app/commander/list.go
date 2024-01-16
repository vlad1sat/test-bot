package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) ListCommand(inputMessage *tgbotapi.Message) {
	outputMSgText := "Here all the products: \n\n"
	products := c.service.List()
	for _, p := range products {
		outputMSgText += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMSgText)
	c.bot.Send(msg)
}
