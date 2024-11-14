package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	"jirku.sk/zberatel/db"
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/pkg/middleware"
	"jirku.sk/zberatel/template/components"
	"jirku.sk/zberatel/template/layout"
	"jirku.sk/zberatel/template/page"
)

type collectionService interface {
	Create(ctx context.Context, input model.CollectionInput) (model.Collection, error)
	Update(ctx context.Context, id model.KSUID, input model.CollectionInput) (model.Collection, error)
	List(ctx context.Context, paging model.PagingRequest) (model.PagingResponse[model.Collection], error)
	Get(ctx context.Context, id model.KSUID) (model.Collection, error)
	Delete(ctx context.Context, id model.KSUID) error
}

type Collection struct {
	log               *slog.Logger
	decoder           *form.Decoder
	queries           *db.Queries
	collectionService collectionService
}

func NewCollection(log *slog.Logger, queries *db.Queries, collectionService collectionService) *Collection {
	return &Collection{
		log:               log,
		decoder:           form.NewDecoder(),
		collectionService: collectionService,
	}
}

func (h *Collection) New(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := h.getViewModel("New Collection", r)
		logger.Debug("sending response", slog.Int("code", http.StatusOK), slog.Any("view-model", viewModel))
		if err := tmpl.ExecuteTemplate(w, "page", viewModel); err != nil {
			slog.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func (h *Collection) getViewModel(title string, r *http.Request) layout.PageVM[page.CollectionsNewModel] {
	csrf := nosurf.Token(r)
	content := page.CollectionsNewModel{
		Form: model.CollectionInput{
			Title: "",
		},
		CsrfToken: csrf,
		TypeOptions: []components.Option{
			{Value: "PopHead", Label: "Pop Head"},
			{Value: "Stamp", Label: "Stamp"},
			{Value: "HotWheels", Label: "Hot Wheels"},
			{Value: "LOL", Label: "Stamp"},
			{Value: "Coins", Label: "Coins"},
		},
	}
	viewModel := layout.NewPageVM(title, csrf, content, r)
	return viewModel
}

func (h *Collection) NewAction(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := h.getViewModel("New Collection", r)

		if err := r.ParseForm(); err != nil {
			logger.Error("parsing form", slog.Any("error", err))
			viewModel.Content.GlobalError = errors.New("could not parse form")
		} else if err := h.decoder.Decode(&viewModel.Content.Form, r.PostForm); err != nil {
			viewModel.Content.GlobalError = errors.New("could not decode form")
		}
		result, err := h.collectionService.Create(r.Context(), viewModel.Content.Form)
		if err == nil {
			logger.Debug("sending response", slog.Int("code", http.StatusOK), slog.Any("view-model", viewModel))
			http.Redirect(w, r, fmt.Sprintf("/collection/%s", result.ID), http.StatusFound)
			return
		}
		logger.Debug("sending response", slog.Int("code", http.StatusBadRequest), slog.Any("view-model", viewModel))
		if err := tmpl.ExecuteTemplate(w, "page", viewModel); err != nil {
			slog.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusBadRequest)
		}
	}
}

func (h *Collection) List(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context(), h.log)
		viewModel := h.getListViewModel("Collections", r)
		ctx := r.Context()
		collections, err := h.queries.GetCollectionsList(ctx, db.GetCollectionsListParams{
			Offset: int32(viewModel.Content.Paging.Start),
			Limit:  int32(viewModel.Content.Paging.Take),
			UserID: viewModel.User.ID,
		})
		if err != nil {
			logger.Error("Failed to laod collections", slog.Any("error", err))
			viewModel.Content.Error = err
		} else {
			items := make([]model.Collection, 0, len(collections))
			for _, c := range collections {
				items = append(items, model.Collection{
					ID:          c.ID,
					Title:       c.Title,
					Type:        c.Type,
					Description: c.Description.String,
				})
			}
			viewModel.Content.Collections.Items = items
		}
	}
}

func (h *Collection) getListViewModel(title string, r *http.Request) layout.PageVM[page.CollectionsListModel] {
	csrf := nosurf.Token(r)
	paging := model.PagingRequest{
		Start: 0,
		Take:  10,
	}
	paging.FromQuery(r.URL.Query())
	content := page.CollectionsListModel{}
	viewModel := layout.NewPageVM(title, csrf, content, r)
	return viewModel
}
