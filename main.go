package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/handler"
)

func main() {
	r := mux.NewRouter()
	r.Use(nosurf.NewPure)
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	auth := handler.NewAuth()
	r.HandleFunc("/auth/login", auth.Login).Methods("GET")
	r.HandleFunc("/auth/register", auth.Register).Methods("GET", "POST")
	err := http.ListenAndServe("localhost:3000", r)
	if err != nil {
		panic(err)
	}
}
