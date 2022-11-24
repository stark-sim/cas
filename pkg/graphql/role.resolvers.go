package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/tools"
	"context"
)

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.CreateRoleInput) (*ent.Role, error) {
	return r.client.Role.Create().SetInput(input).Save(ctx)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id *string, input ent.UpdateRoleInput) (*ent.Role, error) {
	if id != nil {
		tempID := tools.StringToInt64(*id)
		return r.client.Role.UpdateOneID(tempID).SetInput(input).Save(ctx)
	} else {
		return nil, nil
	}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
