package e2e

import (
	"testing"

	"github.com/kelseyhightower/envconfig"
)

type TestConfig struct {
	// APIListenAddress is the IP address and port where the API server listens at.
	APIListenAddress string `default:"127.0.0.1:42069" envconfig:"API_LISTEN_ADDRESS"`
}

func NewTestConfig(t *testing.T) *TestConfig {
	t.Helper()

	var cfg TestConfig

	if err := envconfig.Process("e2e", &cfg); err != nil {
		t.Fatalf("cfg fail to setup: %v", err)
	}

	return &cfg
}
