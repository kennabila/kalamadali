package app

import (
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/kennabila/kalamadali/config"
)

func Kalamadali() http.Handler {
	gotenv.Load()

	route := config.NewRoute()

	return route
}
