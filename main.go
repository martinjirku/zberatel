package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/std"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"jirku.sk/zberatel/handler"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/service"

	_ "github.com/lib/pq"
)

type configuration struct {
	publicPath string
	port       int
	host       string
	logLevel   string
	logHandler string
}

func (c *configuration) Level() slog.Level {
	if c.logLevel == "debug" {
		return slog.LevelDebug
	}
	if c.logLevel == "info" {
		return slog.LevelInfo
	}
	if c.logLevel == "warn" {
		return slog.LevelWarn
	}
	if c.logLevel == "error" {
		return slog.LevelError
	}
	return slog.LevelInfo
}

const (
	GOOGLE_CAPTCHA_ID = "GOOGLE_CAPTCHA_ID"
	GOOGLE_API_KEY    = "GOOGLE_API_KEY"
	DB_ADDR           = "DB_ADDR"
	DB_PORT           = "DB_PORT"
	DB_NAME           = "DB_NAME"
	DB_USER           = "DB_USER"
	DB_PWD            = "DB_PWD"
	SESSION_KEY       = "SESSION_KEY"
)

//go:embed template/**/*.tmpl
var distFS embed.FS

func main() {
	// Channel to signal the server to start
	startChan := make(chan struct{})
	// Handle OS signals.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	config, log := configure()
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", config.host, config.port),
	}

	go func() {
		userSrv, unSrv, storeSrv, collectionSrv := prepareServices(log)
		router := mux.NewRouter()
		setupMiddleware(router, log, storeSrv)
		setupRouter(router, log, userSrv, unSrv, storeSrv, collectionSrv)
		server.Handler = router
		<-startChan
		log.Info(fmt.Sprintf("starting server at %s", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("error starting server", slog.Any("error", err))
			panic(err)
		}
	}()

	select {
	case <-time.After(10 * time.Millisecond): // Adjust timeout as needed
		close(startChan) // Signal to start server
	case <-stopChan:
		log.Info("Startup interrupted, server not started")
		return
	}

	<-stopChan
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Info("Shutting down error", slog.Any("error", err))
	} else {
		log.Info("Server gracefully stopped")
	}
}

func validatePort(port int) bool {
	return port > 0 && port <= 65535
}

func handlerLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "handler"), slog.String("name", name))
}

func middlwareLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "middleware"), slog.String("name", name))
}

func serviceLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "service"), slog.String("name", name))
}

func setupMiddleware(router *mux.Router, log *slog.Logger, store sessions.Store) {
	router.Use(middleware.Recover(middlwareLog(log, "recover")))
	router.Use(middleware.Csrf)
	router.Use(middleware.RequestID(middlwareLog(log, "requestID")))
	router.Use(middleware.AuthMiddleware(store))
	router.Use(middleware.Logger(middlwareLog(log, "logger")))
}

func setupRouter(router *mux.Router, log *slog.Logger, userSrv *service.UserService, unSrv *ut.UniversalTranslator, store sessions.Store, collectionSrv *service.CollectionService) {
	tmpl := getTemplate(log, distFS)
	router.HandleFunc("/", handler.HomeHandler(tmpl("index"))).Methods("GET")
	auth := handler.NewAuth(handlerLog(log, "auth"), os.Getenv(GOOGLE_CAPTCHA_ID), os.Getenv(GOOGLE_API_KEY),
		userSrv, unSrv, store)
	router.HandleFunc("/auth/logout", auth.LogoutAction).Methods("POST")
	router.HandleFunc("/auth/login", auth.Login(tmpl("login"))).Methods("GET")
	router.HandleFunc("/auth/login", auth.LoginAction(tmpl("login"))).Methods("POST")
	router.HandleFunc("/auth/register", auth.Register(tmpl("register"))).Methods("GET")
	router.HandleFunc("/auth/register", auth.RegisterAction(tmpl("register"))).Methods("POST")
	router.HandleFunc("/auth/registration-success", auth.RegistrationSuccess(tmpl("register-success"))).Methods("GET")

	// collection := handler.NewCollection(handlerLog(log, "collection"), collectionSrv)

	collectionRouter := router.PathPrefix("/collections").Subrouter()
	collectionRouter.Use(middleware.AuthorizeMiddleware)
	// collectionRouter.HandleFunc("/new", collection.New).Methods("GET")
	// collectionRouter.HandleFunc("/new", collection.NewAction).Methods("POST")
}

func getTemplate(log *slog.Logger, distFS fs.FS) func(string) *template.Template {
	sproutHandler := sprout.New(
		sprout.WithLogger(log),
		sprout.WithRegistries(
			conversion.NewRegistry(),
			std.NewRegistry(),
			maps.NewRegistry(),
			strings.NewRegistry(),
			slices.NewRegistry(),
		))
	return func(page string) *template.Template {
		path := fmt.Sprintf("template/page/%s.tmpl", page)
		tmpl, err := template.
			New("base").
			Funcs(sproutHandler.Build()).
			ParseFS(distFS, "template/**/*.tmpl", path)
		if err != nil {
			log.Error("loading templates", slog.Any("error", err), slog.String("template", path))
			os.Exit(1)
		}
		return tmpl
	}
}

func prepareServices(log *slog.Logger) (*service.UserService, *ut.UniversalTranslator, sessions.Store, *service.CollectionService) {
	// i18n
	unSrv := ut.New(en.New())
	trans, _ := unSrv.GetTranslator("en")
	validator := validator.New(validator.WithRequiredStructEnabled())
	validator_en.RegisterDefaultTranslations(validator, trans)

	// Database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv(DB_ADDR), os.Getenv(DB_PORT), os.Getenv(DB_USER), os.Getenv(DB_PWD), os.Getenv(DB_NAME))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error("error connecting to database", slog.Any("error", err))
		panic(err)
	}

	// Services
	userSrv := service.NewUserService(serviceLog(log, "userService"), db, validator)

	storeSrv := sessions.NewCookieStore([]byte(os.Getenv(SESSION_KEY)))
	collectionSrv := service.NewCollectionService(log, db, validator)
	return userSrv, unSrv, storeSrv, collectionSrv
}

func configure() (configuration, *slog.Logger) {
	config := configuration{}

	flag.StringVar(&config.publicPath, "public", "", "Usage description of the flag")
	flag.StringVar(&config.host, "host", "localhost", "specify the app host")
	flag.IntVar(&config.port, "port", 3000, "specfiy the port application will listen")
	flag.StringVar(&config.logLevel, "loglevel", "info", "specify the log level (debug, info, warn, error)")
	flag.StringVar(&config.logHandler, "loghandler", "text", "specify the log level (text, json)")
	flag.Parse()

	handlerOptions := &slog.HandlerOptions{
		Level: &config,
	}
	var logHandler slog.Handler
	if config.logHandler == "text" {
		logHandler = slog.NewTextHandler(os.Stdout, handlerOptions)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	}
	log := slog.New(logHandler)

	if !validatePort(config.port) {
		log.Error("invalid port")
		panic("invalid port")
	}

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}
	return config, log
}
