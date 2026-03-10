package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vishal-Choudhary-hi/chotu/internal/handler"
	"github.com/vishal-Choudhary-hi/chotu/internal/middleware"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.ErrorHandler)
	r.Use(middleware.Logging)
	r.Use(middleware.RateLimit)
	r.Use(middleware.AuthMiddleware)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Post("/shorten", handler.ShortenURL)
	r.Get("/{code}", handler.RedirectURL)

	return r
}
