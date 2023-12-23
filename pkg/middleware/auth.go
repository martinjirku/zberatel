package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"jirku.sk/zberatel/model"
)

var Store = sessions.NewCookieStore([]byte("secret-key"))

const SessionName = "session"

type UserCookie int

const UserCookieKey UserCookie = 0

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, err := Store.Get(r, SessionName)
		if err != nil || session.Values["user"] != nil {
			if user, ok := session.Values["user"].(model.UserLogin); ok {
				ctx = context.WithValue(ctx, UserCookieKey, user)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StoreUser(r *http.Request, w http.ResponseWriter, user *model.UserLogin) error {
	session := sessions.NewSession(Store, SessionName)
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
