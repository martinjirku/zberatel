package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"jirku.sk/zberatel/model"
)

const SessionName = "session"

type UserCookie int

const UserCookieKey UserCookie = 0

func AuthMiddleware(store sessions.Store) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			session, err := store.Get(r, SessionName)
			if err != nil || session.Values["user"] != nil {
				if user, ok := session.Values["user"].(model.UserLogin); ok {
					ctx = context.WithValue(ctx, UserCookieKey, user)
				}
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r)
		if user == nil {
			// TODO: handle redirect back to original page
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func StoreUser(r *http.Request, w http.ResponseWriter, user *model.UserLogin, store sessions.Store) error {
	session := sessions.NewSession(store, SessionName)
	if user == nil {
		session.Values["user"] = nil
	} else {
		session.Values["user"] = &user
	}

	session.Options.MaxAge = 60 * 60 * 24 * 10
	session.Options.HttpOnly = true
	session.Options.Path = "/"
	session.Save(r, w)
	return nil
}

func GetUser(r *http.Request) *model.UserLogin {
	result := r.Context().Value(UserCookieKey)
	if result == nil {
		return nil
	} else if result, ok := result.(model.UserLogin); ok {
		return &result
	}
	return nil
}
