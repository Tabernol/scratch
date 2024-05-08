package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"scratch/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("START")

	feed, err := urlToFeed("https://djinni.co/jobs/?primary_keyword=Scala&exp_level=5y")
	if err != nil {
		log.Println("Error caused feeding data", err)
	}
	fmt.Println(feed)

	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	fmt.Printf("PORT: %v", portString)
	fmt.Printf("DB_URL: %v", dbUrl)

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't open connection with DATABASE", err)
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
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
		Addr:    ":" + portString,
	}
	fmt.Println("Server is starting ...")
	servErr := srv.ListenAndServe()
	if servErr != nil {
		log.Fatal("Server problem")
	}
}
