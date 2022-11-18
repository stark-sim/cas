package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"context"
	"fmt"
)

func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Roles(ctx context.Context) ([]*ent.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleResolver) ID(ctx context.Context, obj *ent.Role) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) ID(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDGt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *roleWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) ID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Role returns RoleResolver implementation.
func (r *Resolver) Role() RoleResolver { return &roleResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// RoleWhereInput returns RoleWhereInputResolver implementation.
func (r *Resolver) RoleWhereInput() RoleWhereInputResolver { return &roleWhereInputResolver{r} }

// UserWhereInput returns UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() UserWhereInputResolver { return &userWhereInputResolver{r} }

type queryResolver struct{ *Resolver }
type roleResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type roleWhereInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
