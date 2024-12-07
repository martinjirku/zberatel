package auth

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/golang-jwt/jwt/v5"
	"jirku.sk/kulektor/config"
)

type Role string

const (
	RolePublic    Role = "public"
	RoleCollector Role = "collector"
	RoleAdmin     Role = "admin"
	RoleEditor    Role = "editor"
)

var AllRoles = []Role{
	RolePublic,
	RoleCollector,
	RoleAdmin,
	RoleEditor,
}

func (r Role) String() string {
	return string(r)
}
func (r Role) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, fmt.Sprintf(`"%s"`, r))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (r *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Role must be a string")
	}

	str = strings.ToLower(str) // Ensure case-insensitivity if needed
	role := Role(str)
	if !role.IsValid() {
		return fmt.Errorf("invalid Role: %s", str)
	}

	*r = role
	return nil
}

func (r Role) IsValid() bool {
	for _, validRole := range AllRoles {
		if r == validRole {
			return true
		}
	}
	return false
}

type User struct {
	Email         string
	EmailVerified bool
	Name          string
	Username      string
	Picture       string
	UserID        string
	Roles         []Role
}

func (u User) HasRole(role Role) bool {
	if role == RolePublic && len(u.Roles) == 0 {
		return true
	}
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u User) IsAnonym() bool {
	if len(u.Roles) == 0 {
		return true
	}
	for _, r := range u.Roles {
		if r == RolePublic {
			return true
		}
	}
	return false
}

func (u User) IsCollector() bool {
	for _, r := range u.Roles {
		if r == RoleCollector {
			return true
		}
	}
	return false
}

func (u User) IsAdmin() bool {
	for _, r := range u.Roles {
		if r == RoleAdmin {
			return true
		}
	}
	return false
}

type userContextKey struct{}

var userKey = userContextKey{}

func GetUser(ctx context.Context) User {
	return ctx.Value(userKey).(User)
}

func GetUserAroundOperation(config config.Configuration, log *slog.Logger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		user := User{
			Roles: []Role{RolePublic},
		}
		claims := GetClaims(ctx)
		if GetClaims(ctx) == nil {
			return next(context.WithValue(ctx, userKey, user))
		}
		if claims, ok := (*claims).(jwt.MapClaims); ok {
			if userID, err := claims.GetSubject(); err == nil {
				user.UserID = userID
			}
			if valueRaw, ok := claims["email"]; ok {
				if email, ok := valueRaw.(string); ok {
					user.Email = email
				}
			}
			if valueRaw, ok := claims["email_verified"]; ok {
				if value, ok := valueRaw.(bool); ok {
					user.EmailVerified = value
				}
			}
			if valueRaw, ok := claims["name"]; ok {
				if value, ok := valueRaw.(string); ok {
					user.Name = value
				}
			}
			if valueRaw, ok := claims["nickname"]; ok {
				if value, ok := valueRaw.(string); ok {
					user.Username = value
				}
			}
			if valueRaw, ok := claims["picture"]; ok {
				if value, ok := valueRaw.(string); ok {
					user.Picture = value
				}
			}
			if valueRaw, ok := claims["kulektor/roles"]; ok {
				if value, ok := valueRaw.([]interface{}); ok {
					user.Roles = []Role{}
					for _, v := range value {
						if r, ok := v.(string); ok {
							user.Roles = append(user.Roles, Role(r))
						}
					}
					slog.Debug("claims", slog.Any("roles", user.Roles))
				}
			}
		}
		return next(context.WithValue(ctx, userKey, user))
	}
}
