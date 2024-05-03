package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("START")

	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Printf("PORT: %v", portString)

}
