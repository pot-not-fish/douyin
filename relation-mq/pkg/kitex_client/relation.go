/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 21:35:00
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 21:50:33
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-mq\pkg\kitex_client\relation.go
 */
package kitex_client

import (
	"context"
	"douyin/relation-mq/pkg/kitex_gen/follow_rpc"
	"fmt"
)

var (
	IncFollow int16 = 3

	DecFollow int16 = 4
)

/**
 * @function
 * @description 关注、取关操作
 * @param action_type 操作码 1-关注，2-取消关注
 * @return
 */
func RelationActionRpc(ctx context.Context, action_type int16, user_id, follow_id int64) error {
	respRpc, err := FollowClient.RelationAction(ctx, &follow_rpc.RelationActionReq{
		ActionType: action_type,
		UserId:     user_id,
		FollowId:   follow_id,
	})
	if err != nil {
		return err
	}

	if respRpc.Code != 0 {
		return fmt.Errorf(respRpc.Msg)
	}
	return nil
}
