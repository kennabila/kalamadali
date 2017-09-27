package handler

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	_ "github.com/heroku/x/hmetrics/onload"
)

func kalamadali() {
	bot, err := tgbotapi.NewBotAPI("434434906:AAGFjzj-WBqTdYCueQodvfbguqjIbTL2Fow")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var msg tgbotapi.MessageConfig
	if update.Message == nil {
		continue
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message.Text == "/start" {
		msg = tgbotapi.NewMessage(216200861, "Salam kenal Kakak " + "Ken Nabila" + ", aku Kala")
		msg.ReplyToMessageID = update.Message.MessageID
	} else {
		msg = tgbotapi.NewMessage(216200861, "Hai Kakak " + "Ken Nabila" + ", ada apa?")
		msg.ReplyToMessageID = update.Message.MessageID
	}

	bot.Send(msg)
}
