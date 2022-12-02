package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cas/pkg/ent"
	"cas/pkg/ent/user"
	"cas/pkg/graphql/middlewares"
	"cas/pkg/graphql/model"
	"cas/tools"
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
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
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	if err != nil {
		return nil, err
	}
	return _user, nil
}
