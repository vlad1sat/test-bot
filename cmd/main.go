package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/vlad1sat/bot/internal/app/commander"
	"github.com/vlad1sat/bot/internal/service/product"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	newCommander := commander.NewCommander(bot, productService)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := update.Message

			switch update.Message.Command() {
			case "help":
				newCommander.HelpCommand(msg)
			case "list":
				newCommander.ListCommand(msg)
			default:
				newCommander.DefaultBehavior(msg)
			}
		}
	}
}
