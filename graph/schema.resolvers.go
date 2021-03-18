package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/TheGrizzlyDev/buildhub/graph/generated"
	"github.com/TheGrizzlyDev/buildhub/graph/model"
)

func (r *mutationResolver) CreateBuild(ctx context.Context, build model.InputBuild) (*model.Build, error) {
	id, err := r.BuildRepo.Create(convertInputBuildToEntityBuild(&build))
	if err != nil {
		return nil, err
	}
	return convertInputBuildToOutputBuild(id, &build), nil
}

func (r *mutationResolver) EditBuild(ctx context.Context, id string, build model.InputBuild) (*model.Build, error) {
	err := r.BuildRepo.Edit(id, convertInputBuildToEntityBuild(&build))
	if err != nil {
		return nil, err
	}
	return convertInputBuildToOutputBuild(id, &build), nil
}

func (r *queryResolver) ItemsFor(ctx context.Context, target string) ([]*model.ItemSpec, error) {
	result := []*model.ItemSpec{}
	for _, item := range r.SpecRepo.FindFor(target) {
		convertedItem, err := convertItemSpecToGraphql(item)

		if err != nil {
			return nil, err
		}

		result = append(result, convertedItem)
	}
	return result, nil
}

func (r *queryResolver) Build(ctx context.Context, id string) (*model.Build, error) {
	build, err := r.BuildRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return convertEntityBuildToOutputBuild(id, build), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
