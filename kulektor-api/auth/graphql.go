package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/MicahParks/keyfunc/v3"
	"jirku.sk/kulektor/config"
)

func InitAuth(ctx context.Context, server *handler.Server, log *slog.Logger, conf config.Configuration) {
	// TODO: extract jwks from configuration
	kf, err := keyfunc.NewDefaultCtx(ctx, []string{"https://jirku.eu.auth0.com/.well-known/jwks.json"}) // Context is used to end the refresh goroutine.
	if err != nil {
		log.Error("creating keyfunc", slog.Any("error", err))
	}
	server.AroundOperations(GetClaimsAroundOperation(conf, kf, log))
	server.AroundOperations(GetUserAroundOperation(conf, log))
}

func HasRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (res interface{}, err error) {
	user := GetUser(ctx)
	if user.HasRole(role) {
		return next(ctx)
	}
	return nil, fmt.Errorf("unauthorized, expected role: %s, provided: %s", role, user.Roles)

}
