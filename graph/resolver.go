package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/TheGrizzlyDev/buildhub/items"
	"github.com/TheGrizzlyDev/buildhub/build"
)

type Resolver struct {
	SpecRepo items.ItemSpecRepository
	BuildRepo build.BuildRepository
}

func NewResolver(specRepo *items.EmbedItemSpecRepository, buildRepo *build.InMemoryBuildRepository) (*Resolver) {
	return &Resolver{SpecRepo: specRepo, BuildRepo: buildRepo}
}