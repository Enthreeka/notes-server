package main

import (
	"github.com/Enthreeka/notes-server/internal/app/server"
	"github.com/Enthreeka/notes-server/internal/config"
	"github.com/Enthreeka/notes-server/pkg/logger"
)

func main() {

	configPath := "configs/config.json"

	log := logger.New()

	config, err := config.New(configPath)
	if err != nil {
		log.Fatal("Failed to load config: %s", err)
	}

	server.Run(log, config)
}
