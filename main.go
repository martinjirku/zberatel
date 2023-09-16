package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"jirku.sk/zbera/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}))

	logger.Info("Starting application")
	err := godotenv.Load(".secrets")
	if err != nil {
		logger.Error("Loading '.secrets' file", "error", err)
		os.Exit(1)
	}
	r := router.New(logger)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		logger.Error("Start server", "error", err)
	}
}
