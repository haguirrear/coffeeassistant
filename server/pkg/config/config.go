package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

const (
	localBaseURL = "http://localhost:8000"
	InertiaEntry = "frontend/src/main.tsx"
)

type Config struct {
	IsProd  bool   `env:"IS_PROD"`
	BaseURL string `env:"BASE_URL"`
	Version string `env:"VERSION, default=0.1"`
}

var Conf = MustLoadEnv(context.Background())

func MustLoadEnv(ctx context.Context) Config {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(fmt.Sprintf("error reading config: %v", err))
	}

	if c.BaseURL == "" {
		c.BaseURL = localBaseURL
	}

	return c
}
