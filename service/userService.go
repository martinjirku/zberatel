package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"jirku.sk/zberatel/db"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/pkg/password"
)

type UserService struct {
	log      *slog.Logger
	db       *sql.DB
	queries  *db.Queries
	validate *validator.Validate
}

func NewUserService(log *slog.Logger, sql *sql.DB, validator *validator.Validate) *UserService {
	return &UserService{
		log:      log,
		db:       sql,
		queries:  db.New(sql),
		validate: validator,
	}
}

func (s *UserService) LoginUser(ctx context.Context, input model.UserLoginInput) (model.UserLogin, error) {
	// logger := middleware.GetLogger(ctx, s.log)
	result := model.UserLogin{Username: input.Username}
	resp, err := s.queries.GetUserLogin(ctx, input.Username)
	if err != nil {
		return result, err
	}
	// resp.Password
	ok := password.CheckPasswordHash(input.Password, resp.Password)
	if !ok {
		return result, fmt.Errorf("invalid password")
	}
	result.Email = resp.Email
	result.ID, err = ksuid.FromBytes(resp.ID)
	if err != nil {
		return result, fmt.Errorf("invalid ID")
	}
	return result, nil
}

func (s *UserService) RegisterUser(ctx context.Context, input model.UserRegistrationInput) error {
	logger := middleware.GetLogger(ctx, s.log)
	err := s.validate.Struct(input)
	if err != nil {
		logger.Error("RegisterUser", slog.Any("error", err))
		return err
	}
	logger.Info("RegisterUser", slog.String("username", input.Username), slog.String("email", input.Username))
	ID, err := ksuid.NewRandom()
	if err != nil {
		logger.Error("Generating ID", slog.Any("error", err))
		return fmt.Errorf("error generating ID: %w", err)
	}
	emailToken, err := ksuid.NewRandom()
	if err != nil {
		logger.Error("Generating email token", slog.Any("error", err))
		return fmt.Errorf("error generating email token: %w", err)
	}
	entity := db.RegisterUserParams{
		ID:       ID.Bytes(),
		Username: input.Username,
		Email:    input.Email,
		Token:    emailToken.Bytes(),
	}
	entity.Password, err = password.HashPassword(input.Password)
	if err != nil {
		logger.Error("Hashing password", slog.Any("error", err))
		return fmt.Errorf("error hashing password: %w", err)
	}
	result, err := s.queries.RegisterUser(ctx, entity)
	if err != nil {
		logger.Error("Registering user", slog.Any("error", err))
		return fmt.Errorf("error registering user: %w", err)
	}
	// TODO: finish this later -> email confirmation
	logger.Info("Sending email with", slog.Any("result", result))
	return nil
}
