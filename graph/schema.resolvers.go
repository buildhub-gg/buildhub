package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	buildPackage "github.com/TheGrizzlyDev/buildhub/build"
	"github.com/TheGrizzlyDev/buildhub/graph/generated"
	"github.com/TheGrizzlyDev/buildhub/graph/model"
)

func (r *mutationResolver) CreateBuild(ctx context.Context, build model.InputBuild) (*model.Build, error) {
	id, err := r.BuildRepo.Create(&buildPackage.Build{
		Name: build.Name,
		Items: extractItemsFromInputBuild(build),
	})
	if err != nil {
		return nil, err
	}
	return &model.Build{
		ID:   id,
		Name: build.Name,
	}, nil
}

func (r *mutationResolver) EditBuild(ctx context.Context, id string, build model.InputBuild) (*model.Build, error) {
	err := r.BuildRepo.Edit(id, &buildPackage.Build{
		Name: build.Name,
		Items: extractItemsFromInputBuild(build),
	})
	if err != nil {
		return nil, err
	}
	return &model.Build{
		ID:   id,
		Name: build.Name,
	}, nil
}

func (r *queryResolver) ItemsFor(ctx context.Context, target string) ([]*model.ItemSpec, error) {
	result := []*model.ItemSpec{}
	for _, item := range r.SpecRepo.FindFor(target) {
		attributes := []*model.AttributeSpec{}
		for _, attribute := range item.Attributes {
			attributes = append(attributes, &model.AttributeSpec{
				ID:   attribute.ID,
				Type: model.AttributeType(attribute.Type),
			})
		}
		result = append(result, &model.ItemSpec{
			ID:         item.ID,
			Attributes: attributes,
		})
	}
	return result, nil
}

func (r *queryResolver) Build(ctx context.Context, id string) (*model.Build, error) {
	build, err := r.BuildRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return &model.Build{
		ID:   id,
		Name: build.Name,
	}, nil
}

func extractItemsFromInputBuild(build model.InputBuild) []*buildPackage.Item {
	buildItems := []*buildPackage.Item {}
	for _, item := range build.Items {
		buildItems = append(buildItems, &buildPackage.Item{
			ID: item.ID,
		})
	}
	return buildItems
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
