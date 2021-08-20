package config

import (
	"os"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

func getEnv(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

var Config = struct {
	Database struct {
		Username string
		Password string
		Host     string
		Port     string
		Name     string
		SSLMode  string
	}
}{}

func init() {
	yamlFeeder := &feeder.Yaml{
		Path: getEnv("CONNECT_CONFIG_PATH", "./config.yml"),
	}
	if err := config.New(yamlFeeder).Feed(&Config); err != nil {
		panic(err)
	}

}
