package cmd

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"

	"vanjiii/kook-book-server/internal/api"
	v1 "vanjiii/kook-book-server/internal/api/v1"
	"vanjiii/kook-book-server/internal/recipe"
)

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run application server",
		Long:  "Run application server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			r := chi.NewRouter()

			r.Use(middleware.Logger)

			srv := &http.Server{
				Handler: r,
				Addr:    ":42069",
			}

			recipeService := recipe.NewService(recipe.NewRepository())

			api.RegisterRoutes(
				r,
				*v1.NewRecipe(recipeService),
			)

			var wg sync.WaitGroup

			wg.Add(1)

			doneCh := make(chan struct{})

			go func() {
				defer wg.Done()

				select {
				case <-ctx.Done():
					log.Println("Received shutdown signal, going to shutdown HTTP server")

					ctxTimeout, cancel := context.WithTimeout(ctx, 1)
					defer cancel()

					_ = srv.Shutdown(ctxTimeout) //nolint:contextcheck
				case <-doneCh:
				}
			}()

			defer wg.Wait()
			defer close(doneCh)

			log.Println("Starting HTTP server at :42069")

			err := srv.ListenAndServe()

			// NOTE: ListenAndServe always returns an error.
			log.Println("HTTP server shutdown")
			return err
		},
	}
}
