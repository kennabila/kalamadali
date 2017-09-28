
package handler

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


type Bot struct {
	bot  telegram-bot-api.v4
}

func NewBot() int {
	bot, err := tgbotapi.NewBotAPI("434434906:AAGFjzj-WBqTdYCueQodvfbguqjIbTL2Fow")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot := &Bot{
		bot: bot,
	}

	return bot
}

func (b *Bot) Listen() {

}

func (b *Bot) SendNotification() {
	msg := tgbotapi.NewMessage(216200861, "Salam kenal Kakak "+"Ken Nabila"+", aku Kala")

	b.bot.Send(msg)
}

