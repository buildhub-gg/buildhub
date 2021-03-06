//+build wireinject

package main

import (
	"github.com/TheGrizzlyDev/buildhub/build"
	"github.com/TheGrizzlyDev/buildhub/items"
	"github.com/TheGrizzlyDev/buildhub/graph"
	"github.com/google/wire"
)

func InitResolver() (*graph.Resolver, error) {
	wire.Build(items.NewRepoItemSpecs, graph.NewResolver, build.NewInMemoryBuildRepository)

	return nil, nil
}
