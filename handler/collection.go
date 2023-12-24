package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/segmentio/ksuid"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
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

func (h *Collection) New(w http.ResponseWriter, r *http.Request) {
	vm := layout.NewPageVM("Home", r)
	layout.Page(vm).Render(templ.WithChildren(r.Context(), page.Index()), w)
}
