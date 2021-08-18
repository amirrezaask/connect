package main

import (
	"log"
	"os"

	"github.com/golobby/config/v2"
	"github.com/golobby/config/v2/feeder"
)

var C *config.Config

func getEnv(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func init() {
	var err error
	C, err = config.New(&feeder.Yaml{
		Path: getEnv("CONNECT_CONFIG_PATH", "./config.yml"),
	})
	if err != nil {
		log.Fatalf("cannot create config: %v", err)
	}
}
