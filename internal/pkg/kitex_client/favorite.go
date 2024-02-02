/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-27 11:01:52
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-02 23:39:20
 * @Description:
 * @FilePath: \douyin\internal\pkg\kitex_client\favorite.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"
	"fmt"
)

/**
 * @function
 * @description 点赞视频操作
 * @param action_type 操作码 1-点赞，2-取消点赞
 * @return
 */
func FavoriteActionRpc(ctx context.Context, action_type int16, user_id, video_id int64) (*favorite_rpc.FavoriteActionResp, error) {
	respRpc, err := FavoriteClient.FavoriteAction(ctx, &favorite_rpc.FavoriteActionReq{
		ActionType: action_type,
		UserId:     user_id,
		VideoId:    video_id,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}

/**
 * @function
 * @description 查看一组是是否点赞
 * @param
 * @return
 */
func IsFavoriteRpc(ctx context.Context, userid []int64, videoid []int64) (*favorite_rpc.IsFavoriteResp, error) {
	respRpc, err := FavoriteClient.IsFavorite(ctx, &favorite_rpc.IsFavoriteReq{
		UserId:  userid,
		VideoId: videoid,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}

/**
 * @function
 * @description 查看某个人的点赞视频
 * @param
 * @return
 */
func FavoriteVideo(ctx context.Context, userid, owner_id int64) (*favorite_rpc.FavoriteVideoResp, error) {
	respRpc, err := FavoriteClient.FavoriteVideo(ctx, &favorite_rpc.FavoriteVideoReq{
		UserId:  userid,
		OwnerId: owner_id,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}
