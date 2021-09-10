package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/viper"
)

type Config struct {
	Org   string
	Token string

	Client *tfe.Client
	Ctx    context.Context
}

func New() (*Config, error) {
	// TODO: add more logic for token locations
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	err = validateConfig(&cfg)
	if err != nil {
		return nil, err
	}

	clientConfig := defaultConfig()

	client, err := tfe.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}

	cfg.Client = client
	cfg.Ctx = context.TODO()

	return &cfg, nil
}

// modified from https://github.com/hashicorp/go-tfe/blob/v0.18.0/tfe.go#L67
func defaultConfig() *tfe.Config {
	config := &tfe.Config{
		Address:    "https://app.terraform.io",
		BasePath:   "/api/v2/",
		Token:      "",
		Headers:    make(http.Header),
		HTTPClient: cleanhttp.DefaultPooledClient(),
	}

	// Set the default user agent.
	config.Headers.Set("User-Agent", "tfd")

	return config
}

func validateConfig(conf *Config) error {
	// Token
	if conf.Token == "" {
		return fmt.Errorf("TFD_TOKEN must be set")
	}

	// Org
	if conf.Org == "" {
		return fmt.Errorf("TFD_ORG must be set")
	}

	return nil
}
