package cmd

import (
	"net/http"

	"vanjiii/kook-book-server/internal/api"
	v1 "vanjiii/kook-book-server/internal/api/v1"
	"vanjiii/kook-book-server/internal/recipe"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run application server",
		Long:  "Run application server",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := chi.NewRouter()

			r.Use(middleware.Logger)

			recipeService := recipe.NewService(recipe.NewRepository())

			api.RegisterRoutes(
				r,
				*v1.NewRecipe(recipeService),
			)

			http.ListenAndServe(":3000", r)

			return nil
		},
	}
}
