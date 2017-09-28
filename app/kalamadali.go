package app

import (
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/kennabila/kalamadali/config"
	"github.com/kennabila/kalamadali/handler"
)

type Kalamadali struct {
	Bot     int
	Router  http.Handler
}

func NewKalamadali() *Kalamadali {
	gotenv.Load()

	bot := config.NewBot()
	router := config.NewRoute(bot)

	kalamadali := &Kalamadali{
		Bot: bot,
		Router: router,
	}

	return kalamadali
}
