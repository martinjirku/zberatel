package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	content := page.Login(page.NewLoginVM())
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}

func AuthRegisterHandler(w http.ResponseWriter, r *http.Request) {
	content := page.Register(page.NewRegisterVM())
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}
