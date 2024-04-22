package user

import (
	"context"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
)

type UserServer struct {
}

func NewUserServer() user.UserServer {
	return &UserServer{}
}

func (u *UserServer) Register(ctx context.Context, req *user.RegisterReq) (*user.UserAuthRes, error) {
	//TODO implement me
	panic("Register implement me")
}

func (u *UserServer) Login(ctx context.Context, req *user.LoginReq) (*user.UserAuthRes, error) {
	//TODO implement me
	panic("Login implement me")
}
