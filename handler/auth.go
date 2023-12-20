package handler

import (
	"net/http"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusNotFound)
}
