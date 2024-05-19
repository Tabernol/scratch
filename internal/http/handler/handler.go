package handler

import (
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, "hello Max")
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Something went wrong")
}

//func (apiCfg *api.ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
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
