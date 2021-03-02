package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/TheGrizzlyDev/buildhub/graph/generated"
	"github.com/TheGrizzlyDev/buildhub/graph/model"
)

func (r *mutationResolver) CreateBuild(ctx context.Context, build model.InputBuild) (*model.Build, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditBuild(ctx context.Context, id string, build model.InputBuild) (*model.Build, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemsFor(ctx context.Context, target string) (result []*model.ItemSpec, err error) {
	for _, item := range r.SpecRepo.FindFor(target) {
		result = append(result, &model.ItemSpec{
			ID:         item.ID,
			Attributes: []*model.AttributeSpec{},
		})
	}
	return result, nil
}

func (r *queryResolver) Build(ctx context.Context, id string) (*model.Build, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
