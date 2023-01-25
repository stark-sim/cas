package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stark-sim/cas/pkg/ent"
	"github.com/stark-sim/cas/pkg/ent/role"
	"github.com/stark-sim/cas/pkg/ent/user"
	"github.com/stark-sim/cas/pkg/ent/userrole"
	"github.com/stark-sim/cas/pkg/graphql/middlewares"
	"github.com/stark-sim/cas/pkg/graphql/model"
	"github.com/stark-sim/cas/tools"
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
			return r.client.User.Create().SetName(req.Name).SetPhone(req.Phone).Save(ctx)
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

// CreatedBy is the resolver for the createdBy field.
func (r *roleResolver) CreatedBy(ctx context.Context, obj *ent.Role) (string, error) {
	if obj.CreatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.CreatedBy, 10), nil
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roleResolver) UpdatedBy(ctx context.Context, obj *ent.Role) (string, error) {
	if obj.UpdatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.UpdatedBy, 10), nil
	}
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *userResolver) CreatedBy(ctx context.Context, obj *ent.User) (string, error) {
	if obj.CreatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.CreatedBy, 10), nil
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *userResolver) UpdatedBy(ctx context.Context, obj *ent.User) (string, error) {
	if obj.UpdatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.UpdatedBy, 10), nil
	}
}

// ID is the resolver for the id field.
func (r *userRoleResolver) ID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *userRoleResolver) CreatedBy(ctx context.Context, obj *ent.UserRole) (string, error) {
	if obj.CreatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.CreatedBy, 10), nil
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *userRoleResolver) UpdatedBy(ctx context.Context, obj *ent.UserRole) (string, error) {
	if obj.UpdatedBy == 0 {
		return "", nil
	} else {
		return strconv.FormatInt(obj.UpdatedBy, 10), nil
	}
}

// UserID is the resolver for the userID field.
func (r *userRoleResolver) UserID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// RoleID is the resolver for the roleID field.
func (r *userRoleResolver) RoleID(ctx context.Context, obj *ent.UserRole) (string, error) {
	return strconv.FormatInt(obj.RoleID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createRoleInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateRoleInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createRoleInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateRoleInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDs is the resolver for the userIDs field.
func (r *createRoleInputResolver) UserIDs(ctx context.Context, obj *ent.CreateRoleInput, data []string) error {
	for _, v := range data {
		obj.UserIDs = append(obj.UserIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createUserInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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

// CreatedBy is the resolver for the createdBy field.
func (r *roleWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *roleWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *roleWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *roleWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *roleWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *roleWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *roleWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *roleWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roleWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *roleWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *roleWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *roleWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.RoleWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *roleWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *roleWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *roleWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *roleWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.RoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateRoleInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateRoleInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateRoleInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateRoleInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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

// CreatedBy is the resolver for the createdBy field.
func (r *updateUserInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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

// CreatedBy is the resolver for the createdBy field.
func (r *userRoleWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *userRoleWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *userRoleWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *userRoleWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *userRoleWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *userRoleWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *userRoleWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *userRoleWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *userRoleWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *userRoleWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *userRoleWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *userRoleWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.UserRoleWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *userRoleWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *userRoleWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *userRoleWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *userRoleWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.UserRoleWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
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

// CreatedBy is the resolver for the createdBy field.
func (r *userWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *userWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *userWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *userWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *userWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *userWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *userWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *userWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *userWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *userWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *userWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *userWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *userWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *userWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *userWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *userWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
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
