package graphql

import (
	"context"
	"testing"

	"vanjiii/kook-book-server/pkg/services"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	env, err := services.NewEnv()
	if err != nil {
		t.Fata
	}

	RunGraphqlTest(t, &GraphqlTestEnv{
		Context:  ctx,
		Resolver: NewResolver(env),
		Query:    `query{GetUser(){ID}}`,
		Result:   `"getUser": User{1}`,
		Errors:   []string{},
	})
}
