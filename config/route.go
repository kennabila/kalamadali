package config

import (
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	"github.com/bukalapak/kalamadali/handler"
	midd "github.com/bukalapak/now-you-see-me/middleware"
)

// NewRoute builds all routes needed in Now You See Me
// it also handles CORS configuration
// it is a niladic function that return http.Handler
func NewRoute() http.Handler {
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

	router.Route("/notif", func(r chi.Router) {
		r.Route("/:username", func(r chi.Router) {
			r.Get("/", midd.Monitor(handler.FVT))
		})
	})

	return router
}
