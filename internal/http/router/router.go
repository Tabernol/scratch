package router

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"scratch/internal/database"
	"scratch/internal/http/handler"
)

// APIConfig holds API-related configuration.
type ApiConfig struct {
	DB *database.Queries
}

// NewApiConfig initializes the ApiConfig with a database connection.
func NewApiConfig(db *sql.DB) (*ApiConfig, error) {
	return &ApiConfig{
		DB: database.New(db),
	}, nil
}

//// NewRouter sets up and returns a new router instance.
//func NewRouter() http.Handler {
//	router := chi.NewRouter()
//
//	router.Use(cors.Handler(cors.Options{
//		AllowedOrigins:   []string{"https://*", "http://*"},
//		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//		AllowedHeaders:   []string{"*"},
//		ExposedHeaders:   []string{"Link"},
//		AllowCredentials: false,
//		MaxAge:           300,
//	}))
//
//	v1Router := chi.NewRouter()
//	v1Router.Get("/users", handler.HandleGet)
//	//v1Router.Get("/err", apiConfig.HandleError)
//	//v1Router.Post("/users", apiCfg)
//
//	router.Mount("/v1", v1Router)
//
//	return router
//}

func SetupRouter(db *ApiConfig) http.Handler {
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
	v1Router.Get("/users", handler.HandleGet)
	//v1Router.Post("/users", apiCfg.)

	router.Mount("/v1", v1Router)

	return router
}

//func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
//	type parameters struct {
//		Name string `json:"name"`
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	params := parameters{}
//	err := decoder.Decode(&params)
//	log.Println("Before parsing")
//	if err != nil {
//		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
//		return
//	}
//	log.Println("After parsing")
//
//	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
//		CreatedAt: time.Now().UTC(),
//		UpdatedAt: time.Now().UTC(),
//		Name:      params.Name,
//	})
//	if err != nil {
//		respondWithError(w, 400, fmt.Sprintf("Can't create user: %v", err))
//		return
//	}
//
//	respondWithJson(w, 200, user)
//}
