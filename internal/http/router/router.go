package router

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"scratch/internal/database"
	"scratch/internal/http/handler"
	"scratch/internal/service"
)

func SetupRouter(db *sql.DB) http.Handler {
	router := chi.NewRouter()
	userHandler, err := initService(db)
	if err != nil {
		log.Println("problem in service layer")
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/users", handler.HandleGet)
	v1Router.Post("/users", userHandler.CreateUser)

	router.Mount("/v1", v1Router)

	return router
}

func initService(db *sql.DB) (handler.UserHandler, error) {
	queries := database.New(db)
	userService := &service.UserService{Queries: queries}
	userHandler := &handler.UserHandler{UserService: userService}
	return *userHandler, nil
}
