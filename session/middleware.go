package session

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/segmentio/ksuid"
)

type MiddlewareOpts func(*Middleware)

func NewMiddleware(next http.Handler, opts ...MiddlewareOpts) http.Handler {
	mw := Middleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	for _, opt := range opts {
		opt(&mw)
	}
	return mw
}

func WithSecure(secure bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.Secure = secure
	}
}

func WithHTTPOnly(httpOnly bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.HTTPOnly = httpOnly
	}
}

func WithSessionManager(sessionManager *scs.SessionManager) MiddlewareOpts {
	return func(m *Middleware) {
		m.SessionManager = sessionManager
	}
}

type Middleware struct {
	Next           http.Handler
	SessionManager *scs.SessionManager
	Secure         bool
	HTTPOnly       bool
}

func ID(r *http.Request) (id string) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return
	}
	return cookie.Value
}

func (mw Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := ID(r)
	if id == "" {
		id = ksuid.New().String()
		mw.SessionManager.Put(r.Context(), "sessionID", id)
	}
	mw.Next.ServeHTTP(w, r)
}
