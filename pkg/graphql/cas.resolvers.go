package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/pkg/ent/role"
	"cas/pkg/ent/user"
	"cas/pkg/ent/userrole"
	"cas/pkg/graphql/middlewares"
	"cas/pkg/graphql/model"
	"cas/tools"
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.CreateRoleInput) (*ent.Role, error) {
	return r.client.Role.Create().SetInput(input).Save(ctx)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input ent.UpdateRoleInput) (*ent.Role, error) {
	tempID := tools.StringToInt64(id)
	return r.client.Role.UpdateOneID(tempID).SetInput(input).Save(ctx)
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*ent.Role, error) {
	tempID := tools.StringToInt64(id)
	return r.client.Role.UpdateOneID(tempID).SetDeletedAt(time.Now()).Save(ctx)
}

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

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, req model.RegisterReq) (*ent.User, error) {
	// 先查看有没有重复的手机号用户存在
	_, err := r.client.User.Query().Where(user.Phone(req.Phone), user.DeletedAtEQ(tools.ZeroTime)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return r.client.User.Create().SetName("").SetPhone(req.Phone).Save(ctx)
		} else {
			logrus.Errorf("err at check existing phone: %v", err)
			return nil, err
		}
	} else {
		return nil, errors.New("phone already used by another user")
	}
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	tempID := tools.StringToInt64(id)
	_user, err := r.client.User.Query().Where(user.ID(tempID), user.DeletedAtEQ(tools.ZeroTime)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_role, err := r.client.Role.Query().Where(role.ID(tempID), role.DeletedAtEQ(tools.ZeroTime)).First(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					userRole, err := r.client.UserRole.Query().Where(userrole.ID(tempID), userrole.DeletedAtEQ(tools.ZeroTime)).First(ctx)
					if err != nil {
						return nil, err
					}
					return userRole, nil
				}
				return nil, err
			}
			return _role, nil
		}
		return nil, err
	}
	return _user, nil
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	res := make([]ent.Noder, 0)
	tempIDs := make([]int64, len(ids))
	for _, v := range ids {
		tempIDs = append(tempIDs, tools.StringToInt64(v))
	}
	// User
	users, err := r.client.User.Query().Where(user.IDIn(tempIDs...), user.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range users {
		res = append(res, v)
	}
	// Role
	roles, err := r.client.Role.Query().Where(role.IDIn(tempIDs...), role.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range roles {
		res = append(res, v)
	}
	// UserRole
	userRoles, err := r.client.UserRole.Query().Where(userrole.IDIn(tempIDs...), userrole.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range userRoles {
		res = append(res, v)
	}
	// return
	return res, nil
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context) ([]*ent.Role, error) {
	return r.client.Role.Query().Where(role.DeletedAtEQ(tools.ZeroTime)).All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().Where(user.DeletedAtEQ(tools.ZeroTime)).All(ctx)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, req model.LoginReq) (*ent.User, error) {
	_user, err := r.client.User.Query().Where(user.PhoneEQ(req.Phone), user.DeletedAtEQ(tools.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("login err: %v", err)
		return nil, err
	}
	// 通过 用户 id 生成新 token
	token, err := tools.GetToken(time.Now(), _user.ID)
	if err != nil {
		logrus.Errorf("get token err: %v", err)
		return nil, err
	}
	// 将 token 包装成一个 cookie 返回
	writer := ctx.Value(middlewares.ResponseWriter).(*middlewares.InjectableResponseWriter)
	writer.Cookie = &http.Cookie{
		Name:       tools.CookieName,
		Value:      url.PathEscape(token),
		Path:       "",
		Domain:     "",
		Expires:    time.Now().Add(tools.AccessTokenExp),
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   http.SameSiteLaxMode,
		Raw:        "",
		Unparsed:   nil,
	}
	if err != nil {
		return nil, err
	}
	return _user, nil
}

// ID is the resolver for the id field.
func (r *roleResolver) ID(ctx context.Context, obj *ent.Role) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ID is the resolver for the id field.
func (r *userRoleResolver) ID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// UserID is the resolver for the userID field.
func (r *userRoleResolver) UserID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// RoleID is the resolver for the roleID field.
func (r *userRoleResolver) RoleID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.RoleID, 10), nil
}

// UserIDs is the resolver for the userIDs field.
func (r *createRoleInputResolver) UserIDs(ctx context.Context, obj *ent.CreateRoleInput, data []string) error {
	for _, v := range data {
		obj.UserIDs = append(obj.UserIDs, tools.StringToInt64(v))
	}
	return nil
}

// RoleIDs is the resolver for the roleIDs field.
func (r *createUserInputResolver) RoleIDs(ctx context.Context, obj *ent.CreateUserInput, data []string) error {
	for _, v := range data {
		obj.RoleIDs = append(obj.RoleIDs, tools.StringToInt64(v))
	}
	return nil
}

// ID is the resolver for the id field.
func (r *roleWhereInputResolver) ID(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
	}
	return nil
}

// IDNeq is the resolver for the idNEQ field.
func (r *roleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
	}
	return nil
}

// IDIn is the resolver for the idIn field.
func (r *roleWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *roleWhereInputResolver) IDGt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
	}
	return nil
}

// IDGte is the resolver for the idGTE field.
func (r *roleWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
	}
	return nil
}

// IDLt is the resolver for the idLT field.
func (r *roleWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
	}
	return nil
}

// IDLte is the resolver for the idLTE field.
func (r *roleWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
	}
	return nil
}

// AddUserIDs is the resolver for the addUserIDs field.
func (r *updateRoleInputResolver) AddUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	for _, v := range data {
		obj.AddUserIDs = append(obj.AddUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveUserIDs is the resolver for the removeUserIDs field.
func (r *updateRoleInputResolver) RemoveUserIDs(ctx context.Context, obj *ent.UpdateRoleInput, data []string) error {
	for _, v := range data {
		obj.RemoveUserIDs = append(obj.RemoveUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// AddRoleIDs is the resolver for the addRoleIDs field.
func (r *updateUserInputResolver) AddRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	for _, v := range data {
		obj.AddRoleIDs = append(obj.AddRoleIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveRoleIDs is the resolver for the removeRoleIDs field.
func (r *updateUserInputResolver) RemoveRoleIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	for _, v := range data {
		obj.RemoveRoleIDs = append(obj.RemoveRoleIDs, tools.StringToInt64(v))
	}
	return nil
}

// ID is the resolver for the id field.
func (r *userRoleWhereInputResolver) ID(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
	}
	return nil
}

// IDNeq is the resolver for the idNEQ field.
func (r *userRoleWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
	}
	return nil
}

// IDIn is the resolver for the idIn field.
func (r *userRoleWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userRoleWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *userRoleWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
	}
	return nil
}

// IDGte is the resolver for the idGTE field.
func (r *userRoleWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
	}
	return nil
}

// IDLt is the resolver for the idLT field.
func (r *userRoleWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
	}
	return nil
}

// IDLte is the resolver for the idLTE field.
func (r *userRoleWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
	}
	return nil
}

// ID is the resolver for the id field.
func (r *userWhereInputResolver) ID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
	}
	return nil
}

// IDNeq is the resolver for the idNEQ field.
func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
	}
	return nil
}

// IDIn is the resolver for the idIn field.
func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
	}
	return nil
}

// IDGte is the resolver for the idGTE field.
func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
	}
	return nil
}

// IDLt is the resolver for the idLT field.
func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
	}
	return nil
}

// IDLte is the resolver for the idLTE field.
func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
	}
	return nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

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

type mutationResolver struct{ *Resolver }
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
