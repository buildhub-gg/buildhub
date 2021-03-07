package build

import "errors"

type Service struct {
	repo BuildRepository
}

func (s *Service) Create(build *Build) (string, error) {
	err := s.validate(build)
	if err != nil {
		return "", err
	}
	return s.repo.Create(build)
}

func (s *Service) validate(build *Build) error {
	if build.Name == "" {
		return errors.New("build name cannot be null")
	}
	return nil
}

func NewService(repo BuildRepository) *Service {
	return &Service{
		repo: repo,
	}
}