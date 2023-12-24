package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	if user == nil {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}
	vm := layout.NewPageVM("Home", r)
	layout.Page(vm).Render(templ.WithChildren(r.Context(), page.Index()), w)
}
