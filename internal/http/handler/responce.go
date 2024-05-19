package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Println("REsponding with 5XX error: ", msg)
	}
	type errResponse struct {
		Error string `json"error"`
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON")
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
