package config

import (
	"fmt"
	"os"
)

const (
	YandexAPIKeyEnv = "YA_API_KEY"
	YandexURLEnv    = "YA_API_URL"

	GRPCPortEnv = "GRPC_PORT"
)

var EmptyEnvError = func(name string) error { return fmt.Errorf(fmt.Sprintf("env %s not found", name)) }

type YandexConfig struct {
	ApiKey string
	URL    string
}

type GRPCConfig struct {
	Port string
}

type Config struct {
	YandexConfig YandexConfig
	GRPCConfig   GRPCConfig
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cfg.setEnv(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) setEnv() error {
	val, ok := os.LookupEnv(YandexAPIKeyEnv)
	if !ok {
		return EmptyEnvError(YandexAPIKeyEnv)
	}
	c.YandexConfig.ApiKey = val

	val, ok = os.LookupEnv(YandexURLEnv)
	if !ok {
		return EmptyEnvError(YandexURLEnv)
	}
	c.YandexConfig.URL = val

	val, ok = os.LookupEnv(GRPCPortEnv)
	if !ok {
		return EmptyEnvError(GRPCPortEnv)
	}
	c.GRPCConfig.Port = val

	return nil
}
