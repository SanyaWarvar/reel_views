package app

import (
	"log"
	"rv/config"
	"rv/internal/app/container"

	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	cnt := container.New(cfg)
	if err := cnt.Start(); err != nil {
		log.Printf("Start error: %s", err.Error())
		os.Exit(1)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Start - signal: " + s.String())
	}
	if err := cnt.Stop(); err != nil {
		log.Printf("Stop: %v", err)
	}
}
