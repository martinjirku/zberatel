package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/handler"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/service"
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
	GOOGLE_CAPTCHA_SITE   = "GOOGLE_CAPTCHA_SITE"
	GOOGLE_CAPTCHA_SECRET = "GOOGLE_CAPTCHA_SECRET"
)

func main() {

	config, log := configure()
	userSrv, unSrv := prepareServices(log)

	router := mux.NewRouter()
	setupMiddleware(router, log)
	setupRouter(router, log, userSrv, unSrv)

	// Start server
	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	log.Info(fmt.Sprintf("starting server at %s", addr))
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Error("error starting server", slog.Any("error", err))
		panic(err)
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

func setupMiddleware(router *mux.Router, log *slog.Logger) {
	router.Use(middleware.Recover(middlwareLog(log, "recover")))
	router.Use(nosurf.NewPure)
	router.Use(middleware.RequestID(middlwareLog(log, "requestID")))
	router.Use(middleware.Logger(middlwareLog(log, "logger")))
}

func setupRouter(router *mux.Router, log *slog.Logger, userSrv *service.UserService, unSrv *ut.UniversalTranslator) {
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	auth := handler.NewAuth(
		handlerLog(log, "auth"),
		os.Getenv(GOOGLE_CAPTCHA_SITE),
		os.Getenv(GOOGLE_CAPTCHA_SECRET),
		userSrv,
		unSrv,
	)
	router.HandleFunc("/auth/login", auth.Login).Methods("GET")
	router.HandleFunc("/auth/register", auth.Register).Methods("GET")
	router.HandleFunc("/auth/register", auth.RegisterAction).Methods("POST")
}

func prepareServices(log *slog.Logger) (*service.UserService, *ut.UniversalTranslator) {
	unSrv := ut.New(en.New())
	trans, _ := unSrv.GetTranslator("en")
	validator := validator.New(validator.WithRequiredStructEnabled())
	validator_en.RegisterDefaultTranslations(validator, trans)
	userSrv := service.NewUserService(serviceLog(log, "userService"), validator)
	return userSrv, unSrv
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
