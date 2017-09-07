package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("434434906:AAGFjzj-WBqTdYCueQodvfbguqjIbTL2Fow")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hai Kakak " + update.Message.From.UserName + ", aku Kala")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
