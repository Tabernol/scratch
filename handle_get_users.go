package main

import "net/http"

func handleGet(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}