package handler

import (
	"context"
	"log/slog"

	"github.com/segmentio/ksuid"
	"jirku.sk/zberatel/model"
)

type collectionService interface {
	Create(ctx context.Context, input model.CollectionInput) (model.Collection, error)
	Update(ctx context.Context, id ksuid.KSUID, input model.CollectionInput) (model.Collection, error)
	List(ctx context.Context, paging model.PagingRequest) (model.PagingResponse[model.Collection], error)
	Get(ctx context.Context, id ksuid.KSUID) (model.Collection, error)
	Delete(ctx context.Context, id ksuid.KSUID) error
}

type Collection struct {
	log               *slog.Logger
	collectionService collectionService
}

func NewCollection(log *slog.Logger, collectionService collectionService) *Collection {
	return &Collection{
		log:               log,
		collectionService: collectionService,
	}
}

// func (h *Collection) New(w http.ResponseWriter, r *http.Request) {
// 	content := page.CollectionsNew(page.NewCollectionsNewVM(r))
// 	layout.Page(layout.NewPageVM("Create new Collection", r)).Render(templ.WithChildren(r.Context(), content), w)
// }

// func (h *Collection) NewAction(w http.ResponseWriter, r *http.Request) {
// 	content := page.CollectionsNew(page.NewCollectionsNewVM(r))
// 	layout.Page(layout.NewPageVM("Create new Collection", r)).Render(templ.WithChildren(r.Context(), content), w)
// }
