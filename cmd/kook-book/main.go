package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"vanjiii/kook-book-server/internal/api"
)

func main() {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)

	api.RegisterRoutes(r)

	http.ListenAndServe(":3000", r)
}
