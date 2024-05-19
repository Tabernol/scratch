package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"scratch/config"
	"scratch/internal/database"
	"scratch/internal/http/router"
)

func main() {
	fmt.Println("START")
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("PORT: %v", cfg.Port)
	fmt.Printf("DB_URL: %v", cfg.DatabaseURL)

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error during conecting to database: %v", err)
	}
	defer db.Close()

	r := router.SetupRouter(db)

	log.Printf("Server is running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	// Initialize ApiConfig
	//apiCfg, err := router.SetupRouter(userHandler)
	//if err != nil {
	//	log.Fatalf("Error initializing API config: %v", err)
	//}
	//
	//fmt.Println("Connected to DB")
	//
	//// Set up HTTP routes
	//r := router.SetupRouter(apiCfg)
	//
	//// Start the server
	//log.Printf("Server is running on port %s", cfg.Port)
	//if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}
}
