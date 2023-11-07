package api

import (
	"net/http"

	v1 "vanjiii/kook-book-server/internal/api/v1"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	mux *chi.Mux,

	recipeHandler v1.RecipeHandler,
) {
	mux.Get("/health-check", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/recipe/{recipeID}", recipeHandler.Get)
		r.Post("/recipe", recipeHandler.Create)
	})
}
