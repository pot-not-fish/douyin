/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 21:35:20
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 21:51:13
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-mq\pkg\kitex_client\user.go
 */
package kitex_client

import (
	"context"
	"douyin/relation-mq/pkg/kitex_gen/user_rpc"
	"fmt"
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
