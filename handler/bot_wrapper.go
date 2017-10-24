package handler

import (
	"log"
	"os"
	"strconv"
	"strings"
	"regexp"
  "math/rand"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/kennabila/kalamadali/database"
	"gopkg.in/telegram-bot-api.v4"
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
		DB:  db,
	}

	return botWrapper
}

func (b *BotWrapper) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {
			b.Start(update)
		} else if update.Message.Text == "/stop" {
			b.Stop(update)
		} else if update.Message.Text == "/help" {
			b.Help(update)
		} else if strings.HasPrefix(update.Message.Text, "/update_github") {
			b.UpdateGithub(update)
		}
	}
}

func (b *BotWrapper) Start(u tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	var photo tgbotapi.PhotoConfig

	msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Salam kenal Kakak " + u.Message.From.FirstName + ", aku Kala. Silakan tunggu notifikasi nya yah!")
	msg.ReplyToMessageID = u.Message.MessageID
	b.DB.Insert(strconv.Itoa(u.Message.From.ID), "kennabila")
	photo = tgbotapi.NewPhotoShare(u.Message.Chat.ID, happy_photos[rand.Intn(len(happy_photos)-1)])

	b.Bot.Send(msg)
	b.Bot.Send(photo)
}


func (b *BotWrapper) Stop(u tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	var photo tgbotapi.PhotoConfig

	msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Bye-bye Kakak " + u.Message.From.FirstName + ". Maaf yah aku berisik :(( sampai jumpa lagi nanti!")
	msg.ReplyToMessageID = u.Message.MessageID
	b.DB.Delete(strconv.Itoa(u.Message.From.ID))
	photo = tgbotapi.NewPhotoShare(u.Message.Chat.ID, sad_photos[rand.Intn(len(sad_photos)-1)])

	b.Bot.Send(msg)
	b.Bot.Send(photo)
}


func (b *BotWrapper) UpdateGithub(u tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	var photo tgbotapi.PhotoConfig
	var IsValidGithubId = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	github_id := strings.TrimPrefix(u.Message.Text, "/update_github ")

	if IsValidGithubId(github_id) {
		status := b.DB.Update(strconv.Itoa(u.Message.From.ID), github_id)
		if status == "succeed" {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Berhasil update username github-mu")
		} else {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Tidak bisa update username github, ketik /start untuk memulai")
		}
	} else {
		msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Username github-mu tidak valid, ganti username github dengan mengetikkan /update_github <username_github_mu>")
	}
	msg.ReplyToMessageID = u.Message.MessageID

	b.Bot.Send(msg)
	b.Bot.Send(photo)
}

func (b *BotWrapper) Help(u tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	var photo tgbotapi.PhotoConfig

	msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Ketik /start untuk memulai dan /stop untuk mengakhiri.\nKetik /update_github <username_github_mu> untuk update username github-mu")
	msg.ReplyToMessageID = u.Message.MessageID

	b.Bot.Send(msg)
	b.Bot.Send(photo)
}

func (b *BotWrapper) SendNotification() {
	msg := tgbotapi.NewMessage(216200861, "Salam kenal Kakak "+"Ken Nabila"+", aku Kala")

	b.Bot.Send(msg)
}
