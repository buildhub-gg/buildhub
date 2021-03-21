package build

import "github.com/google/uuid"

type BuildRepository interface {
	Create(build *Build) (string, error)
	Edit(ID string, build *Build) error
	FindById(ID string) (*Build, error)
}

type InMemoryBuildRepository struct {
	data map[string]*Build
}

func (r *InMemoryBuildRepository) Create(build *Build) (string, error) {
	ID := uuid.NewString()
	r.data[ID] = build
	return ID, nil
}

func (r *InMemoryBuildRepository) Edit(ID string, build *Build) error {
	r.data[ID] = build
	return nil
}

func (r *InMemoryBuildRepository) FindById(ID string) (build *Build, err error) {
	return r.data[ID], nil
}

func NewInMemoryBuildRepository() *InMemoryBuildRepository {
	return &InMemoryBuildRepository{
		data: map[string]*Build {},
	}
}