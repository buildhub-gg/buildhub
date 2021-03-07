package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/TheGrizzlyDev/buildhub/item"
	"github.com/TheGrizzlyDev/buildhub/build"
)

type Resolver struct {
	SpecRepo item.ItemSpecRepository
	BuildRepo build.BuildRepository
}

func NewResolver(specRepo *item.EmbedItemSpecRepository, buildRepo *build.InMemoryBuildRepository) (*Resolver) {
	return &Resolver{SpecRepo: specRepo, BuildRepo: buildRepo}
}