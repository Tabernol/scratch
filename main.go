package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"scratch/config"
	"scratch/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("START")
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("PORT: %v", cfg.Port)
	fmt.Printf("DB_URL: %v", cfg.DatabaseURL)

	// Setup database connection
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error setting up database connection: %v", err)
	}
	defer db.Close()

	apiConfig := apiConfig{
		DB: database.New(db),
	}

	// Create a new router
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/users", handleGet)
	v1Router.Get("/err", handleError)
	v1Router.Post("/users", apiConfig.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + cfg.Port,
	}
	fmt.Println("Server is starting ...")
	servErr := srv.ListenAndServe()
	if servErr != nil {
		log.Fatal("Server problem")
	}
}
