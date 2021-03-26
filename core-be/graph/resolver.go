package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/buildhub-gg/buildhub/core-be/item"
	"github.com/buildhub-gg/buildhub/core-be/build"
)

type Resolver struct {
	SpecRepo item.ItemSpecRepository
	BuildRepo build.BuildRepository
}

func NewResolver(specRepo *item.EmbedItemSpecRepository, buildRepo *build.InMemoryBuildRepository) (*Resolver) {
	return &Resolver{SpecRepo: specRepo, BuildRepo: buildRepo}
}