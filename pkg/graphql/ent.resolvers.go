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

func (r *userRoleResolver) ID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleResolver) UserID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleResolver) RoleID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *createRoleInputResolver) UserIDs(ctx context.Context, obj *ent.CreateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createUserInputResolver) RoleIDs(ctx context.Context, obj *ent.CreateUserInput, data []string) error {
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

func (r *updateRoleInputResolver) AddUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateRoleInputResolver) RemoveUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateUserInputResolver) AddRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateUserInputResolver) RemoveRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) ID(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userRoleWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
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

// UserRole returns UserRoleResolver implementation.
func (r *Resolver) UserRole() UserRoleResolver { return &userRoleResolver{r} }

// CreateRoleInput returns CreateRoleInputResolver implementation.
func (r *Resolver) CreateRoleInput() CreateRoleInputResolver { return &createRoleInputResolver{r} }

// CreateUserInput returns CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() CreateUserInputResolver { return &createUserInputResolver{r} }

// RoleWhereInput returns RoleWhereInputResolver implementation.
func (r *Resolver) RoleWhereInput() RoleWhereInputResolver { return &roleWhereInputResolver{r} }

// UpdateRoleInput returns UpdateRoleInputResolver implementation.
func (r *Resolver) UpdateRoleInput() UpdateRoleInputResolver { return &updateRoleInputResolver{r} }

// UpdateUserInput returns UpdateUserInputResolver implementation.
func (r *Resolver) UpdateUserInput() UpdateUserInputResolver { return &updateUserInputResolver{r} }

// UserRoleWhereInput returns UserRoleWhereInputResolver implementation.
func (r *Resolver) UserRoleWhereInput() UserRoleWhereInputResolver {
	return &userRoleWhereInputResolver{r}
}

// UserWhereInput returns UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() UserWhereInputResolver { return &userWhereInputResolver{r} }

type queryResolver struct{ *Resolver }
type roleResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userRoleResolver struct{ *Resolver }
type createRoleInputResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
type roleWhereInputResolver struct{ *Resolver }
type updateRoleInputResolver struct{ *Resolver }
type updateUserInputResolver struct{ *Resolver }
type userRoleWhereInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
