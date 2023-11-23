package api

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"vanjiii/kook-book-server/internal/e2e"

	"github.com/stretchr/testify/assert"
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
			name:     "default",
			endpoint: fmt.Sprintf("http://%s/health-check", cfg.APIListenAddress),
			scode:    200,
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			_, statusCode, err := httpcl.DoRequest(
				ctx,
				tcase.endpoint,
				http.MethodGet,
				nil,
			)

			assert.NoError(t, err)
			assert.Equal(t, tcase.scode, statusCode)
		})
	}
}
