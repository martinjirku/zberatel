package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vm := layout.NewPageVM("Home")
	// _ := middleware.GetUser(r)
	layout.Page(vm).Render(templ.WithChildren(r.Context(), page.Index()), w)
}
