package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"vanjiii/kook-book-server/pkg/auth"
	"vanjiii/kook-book-server/pkg/services"
)

type Resolver struct {
	auth *auth.Handler
}

func NewResolver(env *services.Env) *Resolver {
	return &Resolver{
		auth: env.AuthHandler,
	}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	panic("not implemented")
}

func (r *queryResolver) GetUser(ctx context.Context, ID int) (*User, error) {
	panic("not implemented")
}
