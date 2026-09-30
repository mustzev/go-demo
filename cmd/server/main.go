package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"go-demo/internal/domains/post"
	"go-demo/internal/platform/storage"
	http1 "go-demo/internal/utils/http"
)

const DATABASE_URL = "postgres://postgres:99946632@localhost:5432/demo?sslmode=disable"

func main() {
	db, err := storage.NewDB(DATABASE_URL)
	if err != nil {
		slog.Error("failed to initialize database", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	postRepo := post.NewRepo(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)

	r := chi.NewRouter()
	r.Use(http1.LoggingMiddleware)
	r.Route("/api", func(r chi.Router) {
		r.Mount("/posts", postHandler.Router())
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	slog.Info("server starting", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server stopped", "error", err)
	}
}
