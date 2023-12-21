package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-playground/form"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

type Auth struct {
	decoder  *form.Decoder
	validate *validator.Validate
	ut       *ut.UniversalTranslator
}

func NewAuth() *Auth {
	uni := ut.New(en.New())
	trans, _ := uni.GetTranslator("en")
	validator := validator.New(validator.WithRequiredStructEnabled())
	validator_en.RegisterDefaultTranslations(validator, trans)

	return &Auth{
		decoder:  form.NewDecoder(),
		validate: validator,
		ut:       uni,
	}
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	content := page.Login(page.NewLoginVM())
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	pageVM := page.NewRegisterVM(nosurf.Token(r))

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			pageVM = pageVM.WithMessage("Error parsing form")
		} else if err := h.decoder.Decode(&pageVM.Form, r.PostForm); err != nil {
			pageVM = pageVM.WithMessage("Error decoding form")
		} else {
			err := h.validate.Struct(pageVM.Form)
			if validationErrors := err.(validator.ValidationErrors); validationErrors != nil {
				validationErrors := err.(validator.ValidationErrors)
				for _, e := range validationErrors {
					translator, _ := h.ut.GetTranslator("en")
					fmt.Printf("Error: %s: %s\n", e.Field(), e.Translate(translator))
					if v, ok := pageVM.Form.Errors[e.Field()]; ok {
						pageVM.Form.Errors[e.Field()] = append(v, e.Translate(translator))
					} else {
						pageVM.Form.Errors[e.Field()] = []string{e.Translate(translator)}
					}
				}
			}
		}
	}
	content := page.Register(pageVM)
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}
