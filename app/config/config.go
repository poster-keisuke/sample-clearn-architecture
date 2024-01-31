package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Name     string `envconfig:"DB_NAME" default:"sample"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASSWORD"`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
}

var config Config

func GetConfig() *Config {
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
