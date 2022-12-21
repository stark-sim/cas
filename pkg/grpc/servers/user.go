package servers

import (
	"context"
	"github.com/stark-sim/cas/pkg/ent"
	"github.com/stark-sim/cas/pkg/ent/user"
	"github.com/stark-sim/cas/pkg/grpc/pb"
	"github.com/stark-sim/cas/tools"
)

type UserServer struct {
	Client *ent.Client
}

func (s *UserServer) Get(ctx context.Context, request *__.UserGetRequest) (*__.User, error) {
	_user, err := s.Client.User.Query().Where(user.ID(request.Id), user.DeletedAt(tools.ZeroTime)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &__.User{
		Id:    _user.ID,
		Name:  _user.Name,
		Phone: _user.Phone,
	}, err
}
