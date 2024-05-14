package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"scratch/config"
)

// NewRouter sets up and returns a new router instance.
func NewRouter(apiConfig *config.APIConfig) http.Handler {
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
	v1Router.Get("/users", apiConfig.HandleGet)
	v1Router.Get("/err", apiConfig.HandleError)
	v1Router.Post("/users", apiConfig.HandleCreateUser)

	router.Mount("/v1", v1Router)

	return router
}
