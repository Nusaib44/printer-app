package main

import (
	"log"

	"printer-app/internal/config"
	"printer-app/pkg/api"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Start API server
	log.Println("Starting server...")
	err = api.StartServer(cfg)
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
