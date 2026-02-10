package main

import (
	"esp32-server/internal/config"
	"esp32-server/internal/handlers"
	"esp32-server/internal/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Loading configurations
	cfg := config.Load()

	// 2. Here we will try to connect the database. If we added the database url in config.load()
	// we will be saving incoming data from esp32 to the database. if not we will print out 
	// NO-DATABASE mdoe and continue operation with writing on console.
	var database *storage.DB
	var err error

	if cfg.DatabaseURL != "" {
		database, err = storage.New(cfg.DatabaseURL)
		if err != nil {
			log.Printf("⚠️  Database Connection Failed: %v\n", err)
		} else {
			fmt.Println("✅ Connected to Database")
		}
	} else {
		fmt.Println("ℹ️  Running in NO-DATABASE mode")
	}

	// 3. preparing the database and passing it to the handler so that it can write in it.
	h := &handlers.Handler{
		DB: database,
	}

	// 4. Start the server
	http.HandleFunc("/api/reading", h.IngestData)

	fmt.Println("=================================")
	fmt.Printf("🚀 ESP32 SERVER RUNNING ON %s\n", cfg.Port)
	fmt.Println("=================================")

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}