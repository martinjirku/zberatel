package router

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"jirku.sk/zbera/handlers"
	hasuraauth "jirku.sk/zbera/lib/hasura-auth"
	"jirku.sk/zbera/session"
)

type Options struct {
	HasuraAuthPath           url.URL `json:"hasuraAuthPath"`
	HasuraGraphqlAdminSecret string  `json:"-"`
}

func NewOptions() (Options, error) {
	options := Options{}
	var err error
	if uri := os.Getenv("ZBER_HASURA_AUTH_PATH"); uri == "" {
		err = errors.Join(err, errors.New("check ZBER_HASURA_AUTH_PATH env variable"))
	} else if path, errParse := url.Parse(uri); errParse != nil {
		err = errors.Join(err, errors.New("parse ZBER_HASURA_AUTH_PATH env variable"))
	} else {
		options.HasuraAuthPath = *path
	}

	options.HasuraGraphqlAdminSecret = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if options.HasuraGraphqlAdminSecret == "" {
		err = errors.Join(err, errors.New("check ZBER_HASURA_AUTH_PATH env variable"))
	}
	return options, err
}

func New(l *slog.Logger, options Options) *chi.Mux {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	authService := hasuraauth.NewService(&httpClient, options.HasuraAuthPath, options.HasuraGraphqlAdminSecret)

	r := chi.NewRouter()

	sessions := session.NewSessionManager()

	r.Use(middleware.RequestID)
	// By default set Content-Type to "text/html"
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			h.ServeHTTP(w, r)
		})
	})
	r.Use(requestLogger(l))
	r.Use(middleware.Recoverer)

	// css := home.ClassMaxLoginWidth()
	// style := templ.NewCSSMiddleware(css)

	r.Route("/auth", handlers.NewAuth(l, sessions, authService).Route)
	r.Get("/", handlers.NewHome(l).Index)
	return r
}
