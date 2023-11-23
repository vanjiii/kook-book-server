package e2e

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type client struct {
	cl *http.Client
}

func NewClient() *client {
	return &client{
		cl: &http.Client{},
	}
}

func (c *client) DoRequest(ctx context.Context, addr, method string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, addr, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := c.cl.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}
