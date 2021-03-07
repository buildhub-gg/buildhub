package graph

import (
	buildPackage "github.com/TheGrizzlyDev/buildhub/build"
	"github.com/TheGrizzlyDev/buildhub/graph/model"
)

func convertInputBuildToOutputBuild(ID string, in *model.InputBuild) *model.Build  {
	items := []*model.Item {}

	for _, item := range in.Items {
		items = append(items, &model.Item{
			ID: item.ID,
		})
	}

	return &model.Build{
		ID: ID,
		Name: in.Name,
		Items: items,
	}
}

func convertEntityBuildToOutputBuild(ID string, in *buildPackage.Build) *model.Build {
	items := []*model.Item {}

	for _, item := range in.Items {
		items = append(items, &model.Item{
			ID: item.ID,
		})
	}

	return &model.Build{
		ID: ID,
		Name: in.Name,
		Items: items,
	}
}

func convertInputBuildToEntityBuild(in *model.InputBuild) *buildPackage.Build {
	buildItems := []*buildPackage.Item {}

	for _, item := range in.Items {
		buildItems = append(buildItems, &buildPackage.Item{
			ID: item.ID,
		})
	}

	return &buildPackage.Build{
		Name: in.Name,
		Items: buildItems,
	}
}