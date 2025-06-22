package main

import (
	"log"
	"rv/config"
	"rv/internal/app"
)

const configDir = "./config/main.yaml"

func main() {
	cfg, err := config.NewConfig(configDir)

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run
	app.Run(cfg)
}
