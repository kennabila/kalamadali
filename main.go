package main

import (
	"log"
	"net/http"

	"github.com/kennabila/kalamadali/app"
)

func main() {
	kalamadali := app.NewKalamadali()
	go kalamadali.Bot.Listen()

	log.Println("kalamadali is ready to listen at port 1610")
	http.ListenAndServe(":1610", kalamadali.Router)
}
