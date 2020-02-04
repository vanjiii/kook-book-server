package graphql

import (
	"context"
	"testing"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	env, cleanup := NewTestEnv(t)
	defer cleanup()

	RunGraphqlTest(t, &GraphqlTestEnv{
		Context:  ctx,
		Resolver: NewResolver(env),
		Query:    `mutation{createUser(email: "ivan@dim.dev", password: "1234"){Message}}`,
		Result:   `{"createUser": {"Message": "Success!"}}`,
		Errors:   []string{},
	})
}
