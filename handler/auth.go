package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/a-h/templ"
	"github.com/go-playground/form"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
	"jirku.sk/zberatel/template/partials"
)

type userService interface {
	RegisterUser(ctx context.Context, username, email, password string) error
}

type Auth struct {
	log             *slog.Logger
	decoder         *form.Decoder
	validate        *validator.Validate
	ut              *ut.UniversalTranslator
	recaptchaKey    string
	recaptchaSecret string
	userService     userService
}

func NewAuth(log *slog.Logger, recaptchaKey, recaptchaSecret string, userSrvc userService) *Auth {
	uni := ut.New(en.New())
	trans, _ := uni.GetTranslator("en")
	validator := validator.New(validator.WithRequiredStructEnabled())
	validator_en.RegisterDefaultTranslations(validator, trans)

	return &Auth{
		decoder:         form.NewDecoder(),
		validate:        validator,
		ut:              uni,
		log:             log,
		recaptchaKey:    recaptchaKey,
		recaptchaSecret: recaptchaSecret,
		userService:     userSrvc,
	}
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	content := page.Login(page.NewLoginVM())
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLogger(r.Context(), h.log)
	pageVM := page.NewRegisterVM()
	statusCode := http.StatusOK
	if r.Method == http.MethodPost {
		logger.Info("registering user")
		if err := h.validateCaptcha(r); err != nil {
			pageVM.Message = "Invalid captcha"
			logger.Error("decoding register form data", slog.Any("error", err))
		} else if pageVM.Form, err = h.decodeRegister(r); err != nil {
			pageVM.Message = "Invalid form data"
			logger.Error("decoding register form data", slog.Any("error", err))
		}
		if pageVM.Message == "" {
			err := h.userService.RegisterUser(r.Context(), pageVM.Form.Username, pageVM.Form.Email, pageVM.Form.Password)
			if err != nil {
				pageVM.Message = "Error registering user. Try again later."
			} else {
				// Redirect to success page
				http.Redirect(w, r, "/auth/registration-success", http.StatusFound)
				return
			}
		}
	} else {
		pageVM.Form = partials.NewRegisterFormMV(nosurf.Token(r))
	}

	pageVM.Form.RecaptchaKey = h.recaptchaKey

	logger.Debug("sending response", slog.Int("code", statusCode), slog.Any("view-model", pageVM))
	w.WriteHeader(statusCode)
	content := page.Register(pageVM)
	layout.Page(layout.NewPageVM("Register")).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) decodeRegister(r *http.Request) (partials.RegisterFormMV, error) {
	result := partials.NewRegisterFormMV(nosurf.Token(r))
	if err := r.ParseForm(); err != nil {
		return result, fmt.Errorf("parsing form: %w", err)
	} else if err := h.decoder.Decode(&result, r.PostForm); err != nil {
		return result, fmt.Errorf("decoding formular: %w", err)
	} else if err := h.validate.Struct(result); err != nil {
		if validationErrors := err.(validator.ValidationErrors); validationErrors != nil {
			for _, e := range validationErrors {
				translator, _ := h.ut.GetTranslator("en")
				if v, ok := result.Errors[e.Field()]; ok {
					result.Errors[e.Field()] = append(v, e.Translate(translator))
				} else {
					result.Errors[e.Field()] = []string{e.Translate(translator)}
				}
			}
		}
		return result, fmt.Errorf("validating formular: %w", err)
	}
	return result, nil
}

type googleCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
}

func (h *Auth) validateCaptcha(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("parsing form: %w", err)
	}
	token := r.Form.Get("g-recaptcha-response")
	if token == "" {
		return fmt.Errorf("missing captcha token")
	}
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
		"secret":   {h.recaptchaSecret},
		"response": {token},
	})
	if err != nil {
		return fmt.Errorf("sending captcha request: %w", err)
	}
	defer resp.Body.Close()
	var response googleCaptchaResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("decoding captcha response: %w", err)
	}
	return nil
}
