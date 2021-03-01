package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"github.com/TheGrizzlyDev/buildhub/items"

	"github.com/TheGrizzlyDev/buildhub/graph/generated"
	"github.com/TheGrizzlyDev/buildhub/graph/model"
	yaml "gopkg.in/yaml.v2"
)

func (r *mutationResolver) CreateBuild(ctx context.Context, build model.InputBuild) (*model.Build, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditBuild(ctx context.Context, id string, build model.InputBuild) (*model.Build, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemsFor(ctx context.Context, target string) ([]*model.ItemSpec, error) {
	var specPath string
	var err error
	if specPath, err = filepath.Abs(filepath.Join("items", target+".yml")); err != nil {
		return nil, fmt.Errorf("Cannot find spec for %s", target)
	}
	rawSpec, err := os.ReadFile(specPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Cannot read spec file for %s", target)
	}
	result := []*model.ItemSpec{}
	allItems := bytes.Split(rawSpec, []byte("---"))
	for i, itemRawSpec := range allItems {
		var item *items.ItemSpec
		err = yaml.Unmarshal(itemRawSpec, &item)
		if err != nil {
			return nil, fmt.Errorf("Cannot read spec at index %v", i)
		}
		result = append(result, &model.ItemSpec{
			ID: item.ID,
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
