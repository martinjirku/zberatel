package handler

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/template/layout"
)

func HomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		csrf := nosurf.Token(r)
		model := layout.NewPageVM("Home", csrf, map[string]any{}, r)
		if model.User == nil {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}
		model.Content["username"] = model.User.Username
		if err := tmpl.ExecuteTemplate(w, "page", model); err != nil {
			slog.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}
