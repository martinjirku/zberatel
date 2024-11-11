package layout

import (
	"net/http"

	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
)

type scriptTagType string

var (
	ScriptTypeModule scriptTagType = "module"
	ScriptTypeText   scriptTagType = "text/javascript"
)

type ScriptVM struct {
	Type scriptTagType
	Src  string
}

func NewScriptVM(typ scriptTagType, src string) ScriptVM {
	return ScriptVM{
		Type: typ,
		Src:  src,
	}
}

type PageVM[T any] struct {
	Title     string
	Styles    []string
	Scripts   []ScriptVM
	User      *model.UserLogin
	CsfrToken string
	Content   T
}

func NewPageVM[T any](title, csfr string, content T, r *http.Request) PageVM[T] {
	return PageVM[T]{
		Title:     title,
		Styles:    []string{},
		Scripts:   []ScriptVM{},
		User:      middleware.GetUser(r),
		Content:   content,
		CsfrToken: csfr,
	}
}
