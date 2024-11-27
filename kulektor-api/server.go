package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"
	"jirku.sk/kulektor/auth"
	"jirku.sk/kulektor/config"
	"jirku.sk/kulektor/db"
	"jirku.sk/kulektor/graph"
)

func main() {
	// Channel to signal the server to start
	startChan := make(chan struct{})
	// Handle OS signals.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	conf, log := config.Configure()
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	}
	if conf.Protocol == "https" {
		server.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		pool, err := pgxpool.New(context.Background(), conf.DbURL())
		if err != nil {
			log.Error("Unable to create a connection pool", slog.Any("error", err))
		}
		defer pool.Close()

		mux := http.NewServeMux()
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		queries := db.New(pool)
		resolver := graph.Resolver{
			Queries: queries,
			Pool:    pool,
		}
		gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
			Resolvers: &resolver,
			Directives: graph.DirectiveRoot{
				HasRole: auth.HasRoleDirective,
			},
		}))
		auth.InitAuth(ctx, gqlServer, log, conf)
		gqlServer.Use(extension.FixedComplexityLimit(10000))
		mux.Handle("/query", gqlServer)

		server.Handler = mux
		<-startChan
		log.Info(fmt.Sprintf("starting server at %s", conf.Addr()))
		if conf.Protocol == "https" {
			if err := server.ListenAndServeTLS(conf.SslCertFile, conf.SslKeyFile); err != nil && err != http.ErrServerClosed {
				log.Error("error starting server", slog.Any("error", err))
				panic(err)
			}
		} else {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Error("error starting server", slog.Any("error", err))
				panic(err)
			}
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

	gracefullCtx, cancelGracefull := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelGracefull()

	if err := server.Shutdown(gracefullCtx); err != nil {
		log.Info("Shutting down error", slog.Any("error", err))
	} else {
		log.Info("Server gracefully stopped")
	}
}
