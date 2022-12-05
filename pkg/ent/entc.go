//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("../graphql/gqlgen.yaml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graphql/cas.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	// 整合 versioned migrations 和 GraphQL schema
	if err := entc.Generate("./schema", &gen.Config{
		Features:  []gen.Feature{gen.FeatureVersionedMigration},
		Templates: entgql.AllTemplates,
	}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
