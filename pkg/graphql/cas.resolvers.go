package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"context"
	"fmt"
	"strconv"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context) ([]*ent.Role, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// ID is the resolver for the id field.
func (r *roleResolver) ID(ctx context.Context, obj *ent.Role) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ID is the resolver for the id field.
func (r *userRoleResolver) ID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userID field.
func (r *userRoleResolver) UserID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented: UserID - userID"))
}

// RoleID is the resolver for the roleID field.
func (r *userRoleResolver) RoleID(ctx context.Context, obj *ent.UserRole) (string, error) {
	panic(fmt.Errorf("not implemented: RoleID - roleID"))
}

// UserIDs is the resolver for the userIDs field.
func (r *createRoleInputResolver) UserIDs(ctx context.Context, obj *ent.CreateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented: UserIDs - userIDs"))
}

// RoleIDs is the resolver for the roleIDs field.
func (r *createUserInputResolver) RoleIDs(ctx context.Context, obj *ent.CreateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: RoleIDs - roleIDs"))
}

// ID is the resolver for the id field.
func (r *roleWhereInputResolver) ID(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *roleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *roleWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *roleWhereInputResolver) IDGt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *roleWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *roleWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *roleWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// AddUserIDs is the resolver for the addUserIDs field.
func (r *updateRoleInputResolver) AddUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddUserIDs - addUserIDs"))
}

// RemoveUserIDs is the resolver for the removeUserIDs field.
func (r *updateRoleInputResolver) RemoveUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveUserIDs - removeUserIDs"))
}

// AddRoleIDs is the resolver for the addRoleIDs field.
func (r *updateUserInputResolver) AddRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddRoleIDs - addRoleIDs"))
}

// RemoveRoleIDs is the resolver for the removeRoleIDs field.
func (r *updateUserInputResolver) RemoveRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveRoleIDs - removeRoleIDs"))
}

// ID is the resolver for the id field.
func (r *userRoleWhereInputResolver) ID(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *userRoleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *userRoleWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userRoleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *userRoleWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *userRoleWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *userRoleWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *userRoleWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ID is the resolver for the id field.
func (r *userWhereInputResolver) ID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
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
