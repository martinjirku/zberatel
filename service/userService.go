package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
)

type UserService struct {
	log      *slog.Logger
	db       *sql.DB
	validate *validator.Validate
}

func NewUserService(log *slog.Logger, db *sql.DB, validator *validator.Validate) *UserService {
	return &UserService{
		log:      log,
		db:       db,
		validate: validator,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, input model.UserRegistrationInput) error {
	logger := middleware.GetLogger(ctx, s.log)
	err := s.validate.Struct(input)
	if err != nil {
		logger.Error("RegisterUser", slog.Any("error", err))
		return err
	}
	logger.Info("RegisterUser", slog.String("username", input.Username), slog.String("email", input.Username))
	return nil
}
