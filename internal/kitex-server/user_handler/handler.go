package user_handler

import (
	"context"
	"douyin/internal/pkg/kitex_gen/user_rpc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Userinfo implements the UserServiceImpl interface.
func (u *UserServiceImpl) UserList(ctx context.Context, request *user_rpc.UserListReq) (*user_rpc.UserListResp, error)

func (u *UserServiceImpl) UserAction(ctx context.Context, request *user_rpc.UserActionReq) (*user_rpc.UserActionResp, error)
