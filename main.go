package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"jirku.sk/zberatel/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("POST")
	// r.HandleFunc("/auth/callback", handler.AuthCallbackHandler).Methods("GET")
	http.ListenAndServe("localhost:3000", r)
}
