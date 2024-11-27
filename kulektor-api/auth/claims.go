package auth

import (
	"context"
	"log/slog"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"jirku.sk/kulektor/config"
)

type claimsContextKey struct{}

var claimsKey = claimsContextKey{}

func GetClaims(ctx context.Context) *jwt.Claims {
	if v, ok := ctx.Value(claimsKey).(jwt.Claims); ok {
		return &v
	}
	return nil
}

func GetClaimsAroundOperation(config config.Configuration, kf keyfunc.Keyfunc, log *slog.Logger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		if gqlCtx := graphql.GetOperationContext(ctx); gqlCtx != nil {
			authHeader := gqlCtx.Headers.Get("Authorization")
			if authHeader == "" {
				return next(ctx)
			}
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenStr, kf.KeyfuncCtx(ctx))
			if err != nil {
				return func(ctx context.Context) *graphql.Response {
					return graphql.ErrorResponse(ctx, "invalid token: %s", err)
				}
			}
			return next(context.WithValue(ctx, claimsKey, token.Claims))
		}
		return next(ctx)
	}
}
