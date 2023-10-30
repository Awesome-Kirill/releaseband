package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:8080"`

	HTTPWTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"15s"`
	HTTPRTimeout time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"60s"`
	HTTPITimeout time.Duration `envconfig:"HTTP_IDLE_TIMEOUT" default:"60s"`
}

// New  instance of Config
func New() *Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
