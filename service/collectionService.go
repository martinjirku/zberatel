package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"jirku.sk/zberatel/db"
	"jirku.sk/zberatel/model"
)

type CollectionService struct {
	log      *slog.Logger
	db       *sql.DB
	queries  *db.Queries
	validate *validator.Validate
}

func NewCollectionService(log *slog.Logger, sql *sql.DB, validator *validator.Validate) *CollectionService {
	return &CollectionService{
		log:      log,
		db:       sql,
		queries:  db.New(sql),
		validate: validator,
	}
}

func (srv *CollectionService) Create(ctx context.Context, input model.CollectionInput) (model.Collection, error) {
	return model.Collection{}, nil
}
func (srv *CollectionService) Update(ctx context.Context, id ksuid.KSUID, input model.CollectionInput) (model.Collection, error) {
	return model.Collection{}, nil
}
func (srv *CollectionService) List(ctx context.Context, paging model.PagingRequest) (model.PagingResponse[model.Collection], error) {
	return model.PagingResponse[model.Collection]{
		Total: 0,
		Items: []model.Collection{},
	}, nil
}
func (srv *CollectionService) Get(ctx context.Context, id ksuid.KSUID) (model.Collection, error) {
	return model.Collection{}, nil
}
func (srv *CollectionService) Delete(ctx context.Context, id ksuid.KSUID) error {
	return nil
}
