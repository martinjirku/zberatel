package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

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

	router := mux.NewRouter()
	// Logger middleware
	router.Use(middleware.Recover(middlwareLog(log, "recover")))
	router.Use(nosurf.NewPure)
	router.Use(middleware.RequestID(middlwareLog(log, "requestID")))
	router.Use(middleware.Logger(middlwareLog(log, "logger")))

	// Endpoint handlers
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	userSrv := service.NewUserService(serviceLog(log, "userService"))
	auth := handler.NewAuth(
		handlerLog(log, "auth"),
		os.Getenv(GOOGLE_CAPTCHA_SITE),
		os.Getenv(GOOGLE_CAPTCHA_SECRET),
		userSrv,
	)
	router.HandleFunc("/auth/login", auth.Login).Methods("GET")
	router.HandleFunc("/auth/register", auth.Register).Methods("GET", "POST")

	// Start server
	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	log.Info(fmt.Sprintf("starting server at %s", addr))
	err = http.ListenAndServe(addr, router)
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
