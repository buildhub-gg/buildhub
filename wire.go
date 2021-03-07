//+build wireinject

package main

import (
	"github.com/TheGrizzlyDev/buildhub/build"
	"github.com/TheGrizzlyDev/buildhub/item"
	"github.com/TheGrizzlyDev/buildhub/graph"
	"github.com/google/wire"
)

func InitResolver() (*graph.Resolver, error) {
	wire.Build(item.NewRepoItemSpecs, graph.NewResolver, build.NewInMemoryBuildRepository)

	return nil, nil
}
