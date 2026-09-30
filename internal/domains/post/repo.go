package post

import (
	"context"

	"github.com/uptrace/bun"
)

type repo struct {
	db *bun.DB
}

func NewRepo(db *bun.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) List(ctx context.Context) ([]Post, error) {
	var posts []Post

	err := r.db.NewSelect().
		Model(&posts).
		Order("id DESC").
		Scan(ctx)

	return posts, err
}

func (r *repo) Update(ctx context.Context, id int64, fields UpdatePostRequest) (Post, error) {
	var post Post

	query := r.db.NewUpdate().Model(&post).Where("id = ?", id).Returning("*")

	if fields.Title != nil {
		query = query.Set("title = ?", *fields.Title)
	}

	if fields.Body != nil {
		query = query.Set("body = ?", *fields.Body)
	}

	err := query.Scan(ctx)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}
