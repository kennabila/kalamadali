package handler

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func Kalamadali() {
	bot, err := tgbotapi.NewBotAPI("434434906:AAGFjzj-WBqTdYCueQodvfbguqjIbTL2Fow")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60


	msg := tgbotapi.NewMessage(216200861, "Salam kenal Kakak "+"Ken Nabila"+", aku Kala")

	bot.Send(msg)
}
