package main

import (
	"log"
	"net/http"

	"hu.erettsegizteto/internal/config"
	"hu.erettsegizteto/internal/routers"
	"hu.erettsegizteto/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := storage.NewStorage(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	router := routers.NewRouter(store)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
