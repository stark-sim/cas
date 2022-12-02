package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/pkg/ent/user"
	"cas/tools"
	"context"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(tools.StringToInt64(id)), user.DeletedAtEQ(tools.ZeroTime)).First(ctx)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
