/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-13 10:37:00
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-22 12:58:47
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\kitex-server\user_handler\handler.go
 */
package user_handler

import (
	"context"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/kitex_gen/user_rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Userinfo implements the UserServiceImpl interface.
func (u *UserServiceImpl) UserList(ctx context.Context, request *user_rpc.UserListReq) (*user_rpc.UserListResp, error) {
	resp := new(user_rpc.UserListResp)

	klog.CtxDebugf(ctx, "echo called: UserList")

	userInfoList, err := user_dal.RetreiveUsers(request.UserinfoId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	for _, v := range userInfoList {
		resp.Users = append(resp.Users, &user_rpc.User{
			Id:             int64(v.ID),
			Name:           v.Name,
			FollowCount:    v.FollowCount,
			FollowerCount:  v.FollowerCount,
			Avatar:         v.Avatar,
			Background:     v.Background,
			Signature:      v.Signature,
			TotalFavorited: v.TotalFavorited,
			WorkCount:      v.WorkCount,
			FavoriteCount:  v.FavoriteCount,
		})
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (u *UserServiceImpl) UserAction(ctx context.Context, request *user_rpc.UserActionReq) (*user_rpc.UserActionResp, error) {
	resp := new(user_rpc.UserActionResp)

	klog.CtxDebugf(ctx, "echo called: UserAction")

	user := user_dal.User{
		Name:     request.Username,
		Password: request.Password,
	}

	switch request.ActionType {
	case kitex_client.RegisterUser:
		if err := user.CreateUser(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.LoginUser:
		if err := user.RetrieveAccount(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "invalid action type"
		return resp, nil
	}

	resp.User = &user_rpc.User{
		Id:             int64(user.ID),
		Name:           user.Name,
		FollowCount:    user.FollowCount,
		FollowerCount:  user.FollowerCount,
		Avatar:         user.Avatar,
		Background:     user.Background,
		Signature:      user.Signature,
		TotalFavorited: user.TotalFavorited,
		WorkCount:      user.WorkCount,
		FavoriteCount:  user.FavoriteCount,
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (u *UserServiceImpl) UserInfoAction(ctx context.Context, request *user_rpc.UserInfoActionReq) (*user_rpc.UserInfoActionResp, error) {
	resp := new(user_rpc.UserInfoActionResp)

	klog.CtxDebugf(ctx, "echo called: UserInfoAction")

	switch request.ActionType {
	case kitex_client.IncFavorite:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := user_dal.IncFavorite(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DecFavorite:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := user_dal.DecFavorite(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.IncWorkCount:
		if err := user_dal.IncWorkCount(request.UserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.IncFollow:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := user_dal.IncRelation(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DecFollow:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := user_dal.DecRelation(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "invalid action type"
		return resp, nil
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
