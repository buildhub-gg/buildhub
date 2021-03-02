package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/TheGrizzlyDev/buildhub/items"

type Resolver struct {
	SpecRepo items.ItemSpecRepository
}
