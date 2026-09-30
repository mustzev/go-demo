package post

import (
	"context"
	"errors"
)

type serviceRepo interface {
	List(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, id int64, fields UpdatePostRequest) (Post, error)
}

type service struct {
	repo serviceRepo
}

func NewService(repo serviceRepo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) List(ctx context.Context) ([]Post, error) {
	return s.repo.List(ctx)
}

func (s *service) Update(ctx context.Context, id int64, fields UpdatePostRequest) (Post, error) {
	if fields.Title == nil && fields.Body == nil {
		return Post{}, errors.New("no fields to update")
	}

	return s.repo.Update(ctx, id, fields)
}
