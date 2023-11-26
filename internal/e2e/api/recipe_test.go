package api

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"vanjiii/kook-book-server/internal/e2e"

	"github.com/stretchr/testify/require"
)

func TestApi_Live(t *testing.T) {
	var (
		ctx    = context.TODO()
		cfg    = e2e.NewTestConfig(t)
		httpcl = e2e.NewClient()
	)

	e2e.BootsUpServer(t, httpcl, cfg.APIListenAddress)

	cases := []struct {
		name     string
		endpoint string
		scode    int
	}{
		{
			name:     "get",
			endpoint: fmt.Sprintf("http://%s/api/v1/recipe/%d", cfg.APIListenAddress, 42),
			scode:    http.StatusOK,
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			require := require.New(t)

			resp, statusCode, err := httpcl.DoRequest(
				ctx,
				tcase.endpoint,
				http.MethodGet,
				nil,
			)

			require.NoError(err)
			require.Equal(tcase.scode, statusCode)
			require.Contains(string(resp), "Exampler description")
		})
	}
}
