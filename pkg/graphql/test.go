package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/handler"
)

// GraphqlTestEnv helper struct to hold GQL data to test.
type GraphqlTestEnv struct {
	Context  context.Context
	Resolver *Resolver
	Query    string
	Result   string
	Errors   []string
}

// RunGraphqlTest runs the gql test.
func RunGraphqlTest(t *testing.T, test *GraphqlTestEnv) {
	t.Helper()

	if test.Context != nil {
		test.Context = context.Background()
	}

	schema := handler.GraphQL(NewExecutableSchema(Config{Resolvers: test.Resolver}))

	resp, err := test.exec(t, schema)
	cmpErorrs(t, test.Errors, err)

	got := formatJSON(t, resp)
	want := formatJSON(t, []byte(test.Result))
	if !bytes.Equal(got, want) {
		t.Logf("got:  %s", got)
		t.Logf("want: %s", want)
		t.Fail()
	}
}

// exec executes the actual request
func (g *GraphqlTestEnv) exec(t *testing.T, h http.Handler) ([]byte, []string) {
	t.Helper()

	resp, err := client.New(h).RawPost(g.Query)
	if err != nil {
		t.Fatal(err)
	}

	result, err := json.Marshal(resp.Data)
	if err != nil {
		t.Fatal(err)
	}

	var errs []struct {
		Message string
	}

	errors := make([]string, 0)

	if err := json.Unmarshal(resp.Errors, &errs); err != nil && len(resp.Errors) > 0 {
		t.Fatalf("unmarshal errors: %v", err)
	}
	for _, v := range errs {
		errors = append(errors, v.Message)
	}

	return result, errors
}

func formatJSON(t *testing.T, data []byte) []byte {
	t.Helper()

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil && len(data) > 0 {
		t.Fatalf("invalid JSON: %s; raw data:\n%s", err, data)
	}
	formatted, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return formatted
}

func cmpErorrs(t *testing.T, exp []string, got []string) {
	t.Helper()

	if len(exp) != len(got) {
		t.Errorf("Expected %d errors, got %d: %v", len(exp), len(got), got)
		return
	}

	for i := range got {
		if !strings.Contains(got[i], exp[i]) {
			t.Errorf("Expected error %d to be %q, got %q", i, exp[i], got[i])
		}
	}
}
