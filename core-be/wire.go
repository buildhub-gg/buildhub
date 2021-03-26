//+build wireinject

package main

import (
	"github.com/buildhub-gg/buildhub/core-be/build"
	"github.com/buildhub-gg/buildhub/core-be/item"
	"github.com/buildhub-gg/buildhub/core-be/graph"
	"github.com/google/wire"
)

func InitResolver() (*graph.Resolver, error) {
	wire.Build(item.NewRepoItemSpecs, graph.NewResolver, build.NewInMemoryBuildRepository)

	return nil, nil
}
