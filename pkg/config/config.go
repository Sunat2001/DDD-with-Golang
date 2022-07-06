package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HttpAddr string `default:":8000"`
}

func NewConfig() (*Config, error) {
	var s Config
	err := envconfig.Process("MR", &s)
	if err != nil {
		log.Fatal(err)
	}

	return &s, nil
}
