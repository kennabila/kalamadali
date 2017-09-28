package app

import (
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/kennabila/kalamadali/config"
	"github.com/kennabila/kalamadali/handler"
)

type Kalamadali struct {
	Bot     *handler.BotWrapper
	Router  http.Handler
}

func NewKalamadali() *Kalamadali {
	gotenv.Load()

	bot := handler.NewBotWrapper()
	router := config.NewRoute(bot)

	kalamadali := &Kalamadali{
		Bot: bot,
		Router: router,
	}

	return kalamadali
}
