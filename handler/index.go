package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vm := layout.NewPageVM("Home")
	vm.User = middleware.GetUser(r)
	vm.CsfrToken = nosurf.Token(r)
	layout.Page(vm).Render(templ.WithChildren(r.Context(), page.Index()), w)
}
