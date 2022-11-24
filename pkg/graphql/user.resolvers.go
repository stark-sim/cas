package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/tools"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, input ent.UpdateRoleInput) (*ent.Role, error) {
	if id != nil {
		tempID := tools.StringToInt64(*id)
		return r.client.Role.UpdateOneID(tempID).SetInput(input).Save(ctx)
	} else {
		return nil, nil
	}
}
