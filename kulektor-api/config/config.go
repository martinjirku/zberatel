package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Protocol   string `env:"PROTOCOL,default=http"`
	Host       string `env:"HOST,default=0.0.0.0"`
	Port       int    `env:"PORT,default=8080"`
	LogLevel   string `env:"LOG_LEVEL,default=DEBUG"`
	LogHandler string `env:"LOG_HANDLER,default=text"`
	Db         struct {
		Addr string `env:"DB_ADDR,default=localhost"`
		Port string `env:"DB_PORT,default=5432"`
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pwd  string `env:"DB_PWD"`
	}
	SslCertFile string `env:"SSL_CERT_FILE,default=./certs/cert.pem"`
	SslKeyFile  string `env:"SSL_KEY_FILE,default=./certs/key.pem"`
}

func (c *Configuration) DbURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Db.User, c.Db.Pwd, c.Db.Addr, c.Db.Port, c.Db.Name)
}
func (c *Configuration) Addr() string {
	return fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port)
}

func (c *Configuration) Level() slog.Level {
	if c.LogLevel == "debug" {
		return slog.LevelDebug
	}
	if c.LogLevel == "info" {
		return slog.LevelInfo
	}
	if c.LogLevel == "warn" {
		return slog.LevelWarn
	}
	if c.LogLevel == "error" {
		return slog.LevelError
	}
	return slog.LevelInfo
}
func validatePort(port int) bool {
	return port > 0 && port <= 65535
}

func Configure() (Configuration, *slog.Logger) {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Could not load .env")
	}
	config := Configuration{}

	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	handlerOptions := &slog.HandlerOptions{
		Level: &config,
	}
	var logHandler slog.Handler
	if config.LogHandler == "text" {
		logHandler = slog.NewTextHandler(os.Stdout, handlerOptions)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	}
	log := slog.New(logHandler)

	if !validatePort(config.Port) {
		log.Error("invalid port")
		panic("invalid port")
	}

	return config, log
}
