
package handler

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


type BotWrapper struct {
	Bot  *tgbotapi.BotAPI
}

func NewBotWrapper() *BotWrapper {
	bot, err := tgbotapi.NewBotAPI("434434906:AAGFjzj-WBqTdYCueQodvfbguqjIbTL2Fow")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	botWrapper := &BotWrapper{
		Bot: bot,
	}

	return botWrapper
}

func (b *BotWrapper) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Salam kenal Kakak " + update.Message.From.UserName + ", aku Kala. Silakan tunggu notifikasi nya yah!")
			msg.ReplyToMessageID = update.Message.MessageID
		}

		b.Bot.Send(msg)
	}
}

func (b *BotWrapper) SendNotification() {
	msg := tgbotapi.NewMessage(216200861, "Salam kenal Kakak "+"Ken Nabila"+", aku Kala")

	b.Bot.Send(msg)
}

