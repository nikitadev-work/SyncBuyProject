package main

import (
	"calculation/config"
	"calculation/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Config error: %w", err)
	}

	app.Run(cfg)
}
