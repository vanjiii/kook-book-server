package e2e

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var deadline = time.Now().Add(10 * time.Second)

func BootsUpServer(t *testing.T, cl *client, addr string) {
	t.Helper()

	//nolint:staticcheck
	var err = errors.New("fail to start server")

	for {
		_, code, _ := cl.DoRequest(context.TODO(), fmt.Sprintf("http://%s/health-check", addr), http.MethodGet, nil)

		if code == http.StatusOK {
			err = nil
			break
		}

		if time.Now().After(deadline) {
			err = errors.New("deadline exceeded")
			break
		}
	}

	if err != nil {
		t.Fatal(err)
	}
}
