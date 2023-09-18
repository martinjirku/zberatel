package handlers

import (
	"log/slog"
	"net/http"

	"jirku.sk/zbera/home"
)

type Home struct {
	logger *slog.Logger
}

func NewHome(logger *slog.Logger) *Home {
	return &Home{
		logger: logger,
	}
}

func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	home.Page().Render(r.Context(), w)
}
