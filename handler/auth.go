package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-playground/form"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/partials"
)

type userService interface {
	RegisterUser(ctx context.Context, input model.UserRegistrationInput) error
	LoginUser(ctx context.Context, input model.UserLoginInput) (model.UserLogin, error)
}

type Auth struct {
	log          *slog.Logger
	decoder      *form.Decoder
	ut           *ut.UniversalTranslator
	recaptchaKey string
	googleApiKey string
	userService  userService
	store        sessions.Store
}

func NewAuth(log *slog.Logger, recaptchaKey, googleApiKey string, userSrvc userService, ut *ut.UniversalTranslator, store sessions.Store) *Auth {
	return &Auth{
		decoder:      form.NewDecoder(),
		ut:           ut,
		log:          log,
		recaptchaKey: recaptchaKey,
		googleApiKey: googleApiKey,
		userService:  userSrvc,
		store:        store,
	}
}

func (h *Auth) LogoutAction(w http.ResponseWriter, r *http.Request) {
	middleware.StoreUser(r, w, nil, h.store)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Auth) Login(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUser(r)
		if user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		model := h.getLoginPageModel(r)
		h.renderLogin(w, r, tmpl, model)
	}
}

func (h *Auth) LoginAction(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := h.getLoginPageModel(r)

		if err := h.validateCaptcha(r); err != nil {
			logger.Error("decoding register form data", slog.Any("error", err))
			viewModel.Content.GlobalError = errors.New("invalid captcha")
			h.renderLogin(w, r, tmpl, viewModel)
			return
		}
		if err := h.decodeLoginFormValues(r, &viewModel.Content); err != nil {
			logger.Error("decoding register form data", slog.Any("error", err))
			viewModel.Content.WithGlobalError("Invalid form data")
			h.renderLogin(w, r, tmpl, viewModel)
			return
		}
		userLoginInput := model.NewUserLoginInput(viewModel.Content.Form.Username, viewModel.Content.Form.Password)
		result, err := h.userService.LoginUser(r.Context(), userLoginInput)
		if err != nil {
			viewModel.Content.WithGlobalError("Could not login")
			h.renderLogin(w, r, tmpl, viewModel)
			return
		}
		middleware.StoreUser(r, w, &result, h.store)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *Auth) getLoginPageModel(r *http.Request) layout.PageVM[partials.LoginForm] {
	csrf := nosurf.Token(r)
	loginForm := partials.LoginForm{
		Errors:       map[string]error{},
		RecaptchaKey: h.recaptchaKey,
		Form: partials.LoginFormInput{
			CsfrToken: csrf,
		},
	}

	return layout.NewPageVM("Login", csrf, loginForm, r)
}

func (h *Auth) getRegisterPageModel(r *http.Request) layout.PageVM[partials.RegisterForm] {
	csrf := nosurf.Token(r)
	loginForm := partials.RegisterForm{
		Errors:       map[string][]error{},
		RecaptchaKey: h.recaptchaKey,
		Form: partials.RegisterFormInput{
			CsfrToken: csrf,
		},
	}

	return layout.NewPageVM("Register", csrf, loginForm, r)
}

func (h *Auth) renderLogin(w http.ResponseWriter, r *http.Request, tmpl *template.Template, model layout.PageVM[partials.LoginForm]) {
	if model.Content.GlobalError != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := tmpl.ExecuteTemplate(w, "page", model); err != nil {
		slog.Error("page executing context", slog.Any("error", err))
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}
func (h *Auth) renderRegister(w http.ResponseWriter, r *http.Request, tmpl *template.Template, model layout.PageVM[partials.RegisterForm]) {
	if model.Content.GlobalError != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := tmpl.ExecuteTemplate(w, "page", model); err != nil {
		slog.Error("page executing context", slog.Any("error", err))
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}

func (h *Auth) Register(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUser(r)
		if user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		model := h.getRegisterPageModel(r)
		h.renderRegister(w, r, tmpl, model)
	}
}

func (h *Auth) RegisterAction(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := h.getRegisterPageModel(r)

		if err := h.validateCaptcha(r); err != nil {
			viewModel.Content.GlobalError = errors.New("invalid captcha")
			logger.Error("decoding register form data", slog.Any("error", err))
		} else if err = h.decodeRegisterFormValues(r, &viewModel.Content.Form); err != nil {
			viewModel.Content.GlobalError = errors.New("invalid form data")
			logger.Error("decoding register form data", slog.Any("error", err))
		} else if viewModel.Content.Form.Password != viewModel.Content.Form.PasswordConfirmation {
			logger.Info("checking password confirmation")
			viewModel.Content.SetError("PasswordConfirmation", "Password confirmation does not match the password")
		} else {
			err := h.userService.RegisterUser(r.Context(), model.UserRegistrationInput{
				Username: viewModel.Content.Form.Username,
				Email:    viewModel.Content.Form.Email,
				Password: viewModel.Content.Form.Password,
			})
			if err == nil {
				// Redirect to success page
				http.Redirect(w, r, fmt.Sprintf("/auth/registration-success?username=%s", viewModel.Content.Form.Username), http.StatusFound)
				return
			}
			if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
				translator, _ := h.ut.GetTranslator("en")
				for _, e := range validationErrors {
					viewModel.Content.SetError(e.Field(), e.Translate(translator))
				}
			} else {
				logger.Error("registering user", slog.Any("error", err))
				viewModel.Content.WithGlobalError("Error registering user. Try again later.")
			}
		}
		status := http.StatusOK
		if viewModel.Content.HasError() {
			status = http.StatusBadRequest
		}
		logger.Debug("sending response", slog.Int("code", status), slog.Any("view-model", viewModel))
		w.WriteHeader(status)
		h.renderRegister(w, r, tmpl, viewModel)
	}
}

func (h *Auth) RegistrationSuccess(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := layout.NewPageVM("Successfull Registration", "", map[string]any{
			"username": r.URL.Query().Get("username"),
		}, r)

		logger.Debug("sending response", slog.Int("code", http.StatusOK), slog.Any("view-model", viewModel))
		if err := tmpl.ExecuteTemplate(w, "page", viewModel); err != nil {
			slog.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func (h *Auth) decodeLoginFormValues(r *http.Request, form *partials.LoginForm) error {
	logger := middleware.GetLogger(r.Context(), h.log)

	if err := r.ParseForm(); err != nil {
		logger.Error("parsing form", slog.Any("error", err))
		return fmt.Errorf("parsing form: %w", err)
	}
	if err := h.decoder.Decode(&form.Form, r.PostForm); err != nil {
		logger.Error("decode form", slog.Any("error", err))
		return fmt.Errorf("decoding formular: %w", err)
	}
	return nil
}

func (h *Auth) decodeRegisterFormValues(r *http.Request, form *partials.RegisterFormInput) error {
	logger := middleware.GetLogger(r.Context(), h.log)

	if err := r.ParseForm(); err != nil {
		logger.Error("parsing form", slog.Any("error", err))
		return fmt.Errorf("parsing form: %w", err)
	}
	if err := h.decoder.Decode(&form, r.PostForm); err != nil {
		logger.Error("decode form", slog.Any("error", err))
		return fmt.Errorf("decoding formular: %w", err)
	}
	return nil
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
		"secret":   {h.googleApiKey},
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
