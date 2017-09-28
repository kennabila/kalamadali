package config

import (
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	"github.com/kennabila/kalamadali/handler"
)

// NewRoute builds all routes needed in Now You See Me
// it also handles CORS configuration
// it is a niladic function that return http.Handler
func NewRoute(b *handler.BotWrapper) http.Handler {
	router := chi.NewRouter()

	corsConfig := cors.New(cors.Options{
		// for now, it only accepts GET request
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

	router.Get("/notification", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		b.SendNotification()
	})

	return router
}
