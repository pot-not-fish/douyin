/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-13 19:45:41
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 18:59:01
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_client\user.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/user_rpc"
	"fmt"
)

/**
 * @function
 * @description 查看一组用户信息
 * @param
 * @return
 */
func UserListRpc(ctx context.Context, userid []int64) (*user_rpc.UserListResp, error) {
	respRpc, err := UserinfoClient.UserList(ctx, &user_rpc.UserListReq{
		UserinfoId: userid,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}

var (
	RegisterUser int16 = 8

	LoginUser int16 = 9
)

/**
 * @function
 * @description 注册用户、登录用户
 * @param action_type 操作码 RegisterUser-注册用户 LoginUser-登录用户
 * @return
 */
func UserActionRpc(ctx context.Context, action_type int16, username, password string) (*user_rpc.UserActionResp, error) {
	respRpc, err := UserinfoClient.UserAction(ctx, &user_rpc.UserActionReq{
		ActionType: action_type,
		Username:   username,
		Password:   password,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}

var (
	IncWorkCount int16 = 5
)

/**
 * @function
 * @description 用户信息操作 点赞数自增 点赞数自减 关注数自增 关注数自减
 * @param action_type 操作码 IncUserFavorite-点赞数自增 DecUserFavorite-点赞数自减 IncUserWorkCount-作品数自增 IncUserFollow-关注自增 DecUserFollow-关注自减
 * @return
 */
func UserInfoActionRpc(ctx context.Context, action_type int16, user_id int64, to_user_id *int64) error {
	respRpc, err := UserinfoClient.UserInfoAction(ctx, &user_rpc.UserInfoActionReq{
		ActionType: action_type,
		UserId:     user_id,
		ToUserId:   to_user_id,
	})
	if err != nil {
		return err
	}

	if respRpc.Code != 0 {
		return fmt.Errorf(respRpc.Msg)
	}

	return nil
}
