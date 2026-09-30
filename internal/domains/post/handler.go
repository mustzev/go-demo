package post

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type handlerService interface {
	List(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, id int64, fields UpdatePostRequest) (Post, error)
}

type handler struct {
	service handlerService
}

func NewHandler(service handlerService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.list)
	r.Patch("/{id}", h.update)
	return r
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.service.List(ctx)
	if err != nil {
		http.Error(w, "failed to fetch posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		return
	}
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}

	var body UpdatePostRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	post, err := h.service.Update(ctx, id, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		return
	}
}
