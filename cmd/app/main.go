package main

import (
	"log"
	"rv/config"
	"rv/internal/app"
)

// @title           REEL VIEWS API
// @version         1.0
// @description     This is reel views api service.

const configDir = "./config/main.yaml"

func main() {
	cfg, err := config.NewConfig(configDir)

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run
	app.Run(cfg)
}
