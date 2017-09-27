package app

import (
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/bukalapak/now-you-see-me/config"
)

func Kalamadali() http.Handler {
	gotenv.Load()

	route := config.NewRoute()

	return route
}
