// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/TheGrizzlyDev/buildhub/graph"
	"github.com/TheGrizzlyDev/buildhub/items"
)

// Injectors from wire.go:

func InitResolver() (*graph.Resolver, error) {
	embedItemSpecRepository, err := items.NewRepoItemSpecs()
	if err != nil {
		return nil, err
	}
	resolver := graph.NewResolver(embedItemSpecRepository)
	return resolver, nil
}
