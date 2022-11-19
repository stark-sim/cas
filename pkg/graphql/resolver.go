package graphql

import (
	"cas/pkg/ent"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewExecSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers:  &Resolver{client: client},
		Directives: DirectiveRoot{},
		Complexity: ComplexityRoot{},
	})
}
