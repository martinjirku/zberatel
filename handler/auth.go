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
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
	"jirku.sk/zberatel/template/partials"
)

type userService interface {
	RegisterUser(ctx context.Context, input model.UserRegistrationInput) error
	LoginUser(ctx context.Context, input model.UserLoginInput) (model.UserLogin, error)
}

type Auth struct {
	log             *slog.Logger
	decoder         *form.Decoder
	ut              *ut.UniversalTranslator
	recaptchaKey    string
	recaptchaSecret string
	userService     userService
	store           sessions.Store
}

func NewAuth(log *slog.Logger, recaptchaKey, recaptchaSecret string, userSrvc userService, ut *ut.UniversalTranslator, store sessions.Store) *Auth {
	return &Auth{
		decoder:         form.NewDecoder(),
		ut:              ut,
		log:             log,
		recaptchaKey:    recaptchaKey,
		recaptchaSecret: recaptchaSecret,
		userService:     userSrvc,
		store:           store,
	}
}

func (h *Auth) LogoutAction(w http.ResponseWriter, r *http.Request) {
	middleware.StoreUser(r, w, nil, h.store)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, r, page.NewLoginVM(r, h.recaptchaKey, ""))
}

func (h *Auth) LoginAction(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLogger(r.Context(), h.log)
	pageVM := page.NewLoginVM(r, h.recaptchaKey, "")

	if err := h.validateCaptcha(r); err != nil {
		logger.Error("decoding register form data", slog.Any("error", err))
		h.renderLogin(w, r, pageVM.WithGlobalError("Invalid captcha"))
		return
	}
	if err := h.decodeLoginFormValues(r, &pageVM.Form); err != nil {
		logger.Error("decoding register form data", slog.Any("error", err))
		h.renderLogin(w, r, pageVM.WithGlobalError("Invalid form data"))
		return
	}
	result, err := h.userService.LoginUser(r.Context(), model.NewUserLoginInput(pageVM.Form.Username, pageVM.Form.Password))
	if err != nil {
		h.renderLogin(w, r, pageVM.WithGlobalError("Could not login"))
		return
	}
	middleware.StoreUser(r, w, &result, h.store)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Auth) renderLogin(w http.ResponseWriter, r *http.Request, vm page.LoginVM) {
	if vm.GlobalError != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	content := page.Login(vm)
	layout.Page(layout.NewPageVM("Login", r)).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	pageVM := page.NewRegisterVM(r, h.recaptchaKey)
	content := page.Register(pageVM)
	layout.Page(layout.NewPageVM("Register", r)).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) RegisterAction(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLogger(r.Context(), h.log)
	pageVM := page.NewRegisterVM(r, h.recaptchaKey)

	if err := h.validateCaptcha(r); err != nil {
		pageVM.Message = "Invalid captcha"
		logger.Error("decoding register form data", slog.Any("error", err))
	} else if err = h.decodeRegisterFormValues(r, &pageVM.Form); err != nil {
		pageVM.Message = "Invalid form data"
		logger.Error("decoding register form data", slog.Any("error", err))
	} else if pageVM.Form.Password != pageVM.Form.PasswordConfirmation {
		logger.Info("checking password confirmation")
		pageVM.Form.Errors = map[string][]string{"PasswordConfirmation": {"Password confirmation does not match the password"}}
	} else {
		err := h.userService.RegisterUser(r.Context(), model.UserRegistrationInput{
			Username: pageVM.Form.Username,
			Email:    pageVM.Form.Email,
			Password: pageVM.Form.Password,
		})
		if err == nil {
			// Redirect to success page
			http.Redirect(w, r, "/auth/registration-success", http.StatusFound)
			return
		}
		if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
			translator, _ := h.ut.GetTranslator("en")
			for _, e := range validationErrors {
				pageVM.Form.SetError(e.Field(), e.Translate(translator))
			}
		} else {
			logger.Error("registering user", slog.Any("error", err))
			pageVM.Message = "Error registering user. Try again later."
		}
	}
	status := http.StatusOK
	if pageVM.Message != "" || !pageVM.Form.IsValid() {
		status = http.StatusBadRequest
	}
	logger.Debug("sending response", slog.Int("code", status), slog.Any("view-model", pageVM))
	w.WriteHeader(status)

	content := page.Register(pageVM)
	layout.Page(layout.NewPageVM("Register", r)).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) RegistrationSuccess(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLogger(r.Context(), h.log)
	username := r.URL.Query().Get("username")

	pageVM := page.RegigisterSucessVM{
		Title:    "Registration success",
		Username: username,
	}

	logger.Debug("sending response", slog.Int("code", http.StatusOK), slog.Any("view-model", pageVM))
	content := page.RegisterSuccess(pageVM)
	layout.Page(layout.NewPageVM("Successfull Registration", r)).Render(templ.WithChildren(r.Context(), content), w)
}

func (h *Auth) decodeLoginFormValues(r *http.Request, form *partials.LoginFormVM) error {
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

func (h *Auth) decodeRegisterFormValues(r *http.Request, form *partials.RegisterFormMV) error {
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
