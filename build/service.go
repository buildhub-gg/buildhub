package build

import "errors"

type Service struct {
	repo BuildRepository
}

func (s *Service) Create(build *Build) (string, error) {
	if build.Name == "" {
		return "", errors.New("build name cannot be null")
	}
	return s.repo.Create(build)
}

func NewService(repo BuildRepository) *Service {
	return &Service{
		repo: repo,
	}
}