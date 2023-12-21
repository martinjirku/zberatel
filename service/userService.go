package service

import (
	"context"
	"log/slog"

	"jirku.sk/zberatel/pkg/middleware"
)

type UserService struct {
	log *slog.Logger
}

func NewUserService(log *slog.Logger) *UserService {
	return &UserService{
		log: log,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, username, email, password string) error {
	logger := middleware.GetLogger(ctx, s.log)
	logger.Info("RegisterUser", slog.String("username", username), slog.String("email", email), slog.String("password", password))
	return nil
}
