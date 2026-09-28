package post

import (
	"context"

	"github.com/uptrace/bun"
)

type Repo struct {
	db *bun.DB
}

func NewRepo(db *bun.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) List(ctx context.Context) ([]Post, error) {
	var posts []Post

	err := r.db.NewSelect().
		Model(&posts).
		Order("id DESC").
		Scan(ctx)

	return posts, err
}
