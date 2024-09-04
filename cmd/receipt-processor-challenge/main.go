package main

import (
	"log"
	"net/http"
	"receipt-processor-challenge/internal/storage"
	"receipt-processor-challenge/pkg/api"
	"receipt-processor-challenge/pkg/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize storage
	storage := storage.New()

	// Setup routes with the storage instance
	router := api.SetupRouter(storage)

	// Start the server
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
