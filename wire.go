//+build wireinject

package main

import (
	"github.com/buildhub-gg/buildhub/build"
	"github.com/buildhub-gg/buildhub/item"
	"github.com/buildhub-gg/buildhub/graph"
	"github.com/google/wire"
)

func InitResolver() (*graph.Resolver, error) {
	wire.Build(item.NewRepoItemSpecs, graph.NewResolver, build.NewInMemoryBuildRepository)

	return nil, nil
}
