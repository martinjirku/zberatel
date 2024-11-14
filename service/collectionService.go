package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/go-playground/validator/v10"
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
	collection := db.CreateCollectionParams{
		ID:     model.NewKSUID(),
		UserID: input.UserID,
		Title:  input.Title,
		Description: sql.NullString{
			String: input.Description,
			Valid:  true,
		},
		Type: input.Type,
	}
	c, errCreate := srv.queries.CreateCollection(ctx, collection)
	if errCreate != nil {
		return model.Collection{}, errCreate
	}
	return c.ToModel(), nil
}
func (srv *CollectionService) Update(ctx context.Context, id model.KSUID, input model.CollectionInput) (model.Collection, error) {
	return model.Collection{}, nil
}
func (srv *CollectionService) List(ctx context.Context, paging model.PagingRequest) (model.PagingResponse[model.Collection], error) {
	srv.queries.GetCollectionsList(ctx, db.GetCollectionsListParams{
		Offset: int32(paging.Start),
		Limit:  int32(paging.Take),
	})
	return model.PagingResponse[model.Collection]{
		Total: 0,
		Items: []model.Collection{},
	}, nil
}
func (srv *CollectionService) Get(ctx context.Context, id model.KSUID) (model.Collection, error) {
	collection, err := srv.queries.GetCollectionByID(ctx, id)
	if err != nil {
		return model.Collection{}, nil
	}
	return collection.ToModel(), nil
}
func (srv *CollectionService) Delete(ctx context.Context, id model.KSUID) error {
	return nil
}
