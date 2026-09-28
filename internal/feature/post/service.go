package post

import "context"

type Service struct {
	Repo *Repo
}

func NewService(Repo *Repo) *Service {
	return &Service{
		Repo: Repo,
	}
}

func (s *Service) List(ctx context.Context) ([]Post, error) {
	return s.Repo.List(ctx)
}
