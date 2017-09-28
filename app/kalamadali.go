package app

import (
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/kennabila/kalamadali/config"
	"github.com/kennabila/kalamadali/handler"
	"github.com/kennabila/kalamadali/database"
)

type Kalamadali struct {
	Bot    *handler.BotWrapper
	Router http.Handler
	DB	   *database.Database
}

func NewKalamadali() *Kalamadali {
	gotenv.Load()

	db := database.NewDatabase()
	bot := handler.NewBotWrapper(db)
	router := config.NewRoute(bot)

	kalamadali := &Kalamadali{
		Bot:    bot,
		Router: router,
	}

	return kalamadali
}
