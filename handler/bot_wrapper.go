package handler

import (
	"log"
	"os"
	"strconv"

	_ "github.com/heroku/x/hmetrics/onload"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/kennabila/kalamadali/database"
)

type BotWrapper struct {
	Bot *tgbotapi.BotAPI
	DB  *database.Database
}

func NewBotWrapper(db *database.Database) *BotWrapper {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	botWrapper := &BotWrapper{
		Bot: bot,
		DB: db,
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
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Salam kenal Kakak "+update.Message.From.UserName+", aku Kala. Silakan tunggu notifikasi nya yah!")
			msg.ReplyToMessageID = update.Message.MessageID
			b.DB.Insert(strconv.Itoa(update.Message.From.ID), "kennabila")

		} else if update.Message.Text == "/stop" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Bye-bye Kakak "+update.Message.From.UserName+". Maaf yah aku berisik :') sampai jumpa lagi nanti!")
			msg.ReplyToMessageID = update.Message.MessageID
		} else if update.Message.Text == "/help" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ketik /start untuk memulai dan /stop untuk mengakhiri.\nKetik /update_github untuk update github username")
			msg.ReplyToMessageID = update.Message.MessageID
		} else if update.Message.Text == "/update_github" {
			//cek dia terdaftar apa ngk, kalau ngk suruh start dulu
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Update github username")
			msg.ReplyToMessageID = update.Message.MessageID
		}

		b.Bot.Send(msg)
	}
}

func (b *BotWrapper) SendNotification() {
	msg := tgbotapi.NewMessage(216200861, "Salam kenal Kakak "+"Ken Nabila"+", aku Kala")

	b.Bot.Send(msg)
}
