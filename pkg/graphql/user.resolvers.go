package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/tools"
	"context"
	"time"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input ent.UpdateUserInput) (*ent.User, error) {
	tempID := tools.StringToInt64(id)
	return r.client.User.UpdateOneID(tempID).SetInput(input).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*ent.User, error) {
	tempID := tools.StringToInt64(id)
	return r.client.User.UpdateOneID(tempID).SetDeletedAt(time.Now()).Save(ctx)
}
