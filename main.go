package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"jirku.sk/zberatel/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("GET")
	r.HandleFunc("/auth/register", handler.AuthRegisterHandler).Methods("GET")
	err := http.ListenAndServe("localhost:3000", r)
	if err != nil {
		panic(err)
	}
}
