/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-14 15:23:04
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-02 16:08:11
 * @Description:
 * @FilePath: \douyin\internal\pkg\kitex_client\follow.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/follow_rpc"
	"fmt"
)

/**
 * @function
 * @description 查看是否关注某一组用户
 * @param
 * @return
 */
func IsFollowRpc(ctx context.Context, user_id int64, to_user_id_list []int64) (*follow_rpc.IsFollowResp, error) {
	respRpc, err := FollowClient.IsFollow(ctx, &follow_rpc.IsFollowReq{
		UserId:   user_id,
		FollowId: to_user_id_list,
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
	IncFollow int16 = 1

	DecFollow int16 = 2
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

var (
	FollowList int16 = 1

	FollowerList int16 = 2
)

/**
 * @function
 * @description 查看关注列表、粉丝列表
 * @param action_type 操作码 1-关注列表 2-粉丝列表
 * @return
 */
func RelationListRpc(ctx context.Context, action_type int16, owner_id int64) (*follow_rpc.RelationListResp, error) {
	respRpc, err := FollowClient.RelationList(ctx, &follow_rpc.RelationListReq{
		ActionType: action_type,
		UserId:     owner_id,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}
	return respRpc, nil
}
