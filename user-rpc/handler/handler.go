/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 20:06:35
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 11:30:40
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\user-rpc\handler\handler.go
 */
package handler

import (
	"context"
	dao "douyin/user-rpc/pkg/dao"
	"douyin/user-rpc/user_rpc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Userinfo implements the UserServiceImpl interface.
func (u *UserServiceImpl) UserList(ctx context.Context, request *user_rpc.UserListReq) (*user_rpc.UserListResp, error) {
	resp := new(user_rpc.UserListResp)

	userInfoList, err := dao.RetreiveUsers(request.UserinfoId)
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

var (
	RegisterUser int16 = 8

	LoginUser int16 = 9
)

func (u *UserServiceImpl) UserAction(ctx context.Context, request *user_rpc.UserActionReq) (*user_rpc.UserActionResp, error) {
	resp := new(user_rpc.UserActionResp)

	user := dao.User{
		Name:     request.Username,
		Password: request.Password,
	}

	switch request.ActionType {
	case RegisterUser:
		if err := user.CreateUser(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case LoginUser:
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

var (
	IncFavorite int16 = 1

	DecFavorite int16 = 2

	IncFollow int16 = 3

	DecFollow int16 = 4

	IncWorkCount int16 = 5
)

func (u *UserServiceImpl) UserInfoAction(ctx context.Context, request *user_rpc.UserInfoActionReq) (*user_rpc.UserInfoActionResp, error) {
	resp := new(user_rpc.UserInfoActionResp)

	switch request.ActionType {
	case IncFavorite:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := dao.IncFavorite(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DecFavorite:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := dao.DecFavorite(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case IncWorkCount:
		if err := dao.IncWorkCount(request.UserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case IncFollow:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := dao.IncRelation(request.UserId, *request.ToUserId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DecFollow:
		if request.ToUserId == nil || *request.ToUserId <= 0 {
			resp.Code = 1
			resp.Msg = "invalid to user id"
			return resp, nil
		}

		if err := dao.DecRelation(request.UserId, *request.ToUserId); err != nil {
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
