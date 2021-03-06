package config

import (
	"net/http"
	"time"
	"fmt"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	"github.com/kennabila/kalamadali/handler"
)

func NewRoute(b *handler.BotWrapper) http.Handler {
	router := chi.NewRouter()

	corsConfig := cors.New(cors.Options{
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
	})

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(corsConfig.Handler)

	router.Post("/notification", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("aaaaa")
		w.Write([]byte("welcome"))
		b.SendNotification()
	})

	return router
}
