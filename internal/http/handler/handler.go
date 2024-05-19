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
