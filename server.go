package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TheGrizzlyDev/buildhub/graph/generated"
	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")

	graphResolver, err := InitResolver()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// Source edited from: https://gqlgen.com/recipes/cors/
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// Adder router to ListenAndServe src: https://github.com/99designs/gqlgen/issues/1326#issuecomment-693454584
	log.Fatal(http.ListenAndServe(":"+port, router))
}
