package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"jirku.sk/zbera/auth"
	"jirku.sk/zbera/components"
	hasuraauth "jirku.sk/zbera/lib/hasura-auth"
)

type AuthService interface {
	Login(ctx context.Context, req hasuraauth.RequestEmailPassword) (result hasuraauth.ResponseEmailPassword, status int, err error)
}

type Auth struct {
	Log            *slog.Logger
	SessionManager *scs.SessionManager
	AuthService    AuthService
}

func NewAuth(l *slog.Logger, s *scs.SessionManager, a AuthService) *Auth {
	return &Auth{
		Log:            l,
		SessionManager: s,
		AuthService:    a,
	}
}

func (a *Auth) Route(r chi.Router) {
	r.Get("/login", a.Login)
	r.Post("/login", a.LoginPost)
	r.Post("/logout", a.Logout)
	r.Post("/register", a.Register)
	r.Post("/forgot-password", a.ForgotPassword)
	r.Post("/reset-password", a.ResetPassword)
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	components.Layout("Login", auth.Page(components.LoginForm("login"))).Render(r.Context(), w)
}

func (a *Auth) LoginPost(w http.ResponseWriter, r *http.Request) {
	request := hasuraauth.RequestEmailPassword{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	result, status, err := a.AuthService.Login(r.Context(), request)
	if err != nil {
		a.Log.Error("Login", "error", err)
		w.WriteHeader(status)
		return
	}
	ctx := r.Context()
	a.SessionManager.Put(ctx, "session", result.Session)
	w.WriteHeader(status)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}
func (a *Auth) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ForgotPassword"))
}
func (a *Auth) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ResetPassword"))
}
