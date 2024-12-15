package main

import (
	"log"

	"printer-app/pkg/api"
)

func main() {
	// // Load configuration
	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Fatalf("Failed to load config: %v", err)
	// }

	// Start API server
	log.Println("Starting server...")
	err := api.StartServer()
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
