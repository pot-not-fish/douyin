/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-27 11:01:52
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 18:58:39
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_client\favorite.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"
	"fmt"
)

var (
	IncFavorite int16 = 1

	DecFavorite int16 = 2
)

/**
 * @function
 * @description 点赞视频操作
 * @param action_type 操作码 IncFavorite-点赞，DecFavorite-取消点赞
 * @return
 */
func FavoriteActionRpc(ctx context.Context, action_type int16, user_id, video_id int64) error {
	respRpc, err := FavoriteClient.FavoriteAction(ctx, &favorite_rpc.FavoriteActionReq{
		ActionType: action_type,
		UserId:     user_id,
		VideoId:    video_id,
	})
	if err != nil {
		return err
	}

	if respRpc.Code != 0 {
		return fmt.Errorf(respRpc.Msg)
	}

	return nil
}

/**
 * @function
 * @description 查看一组是是否点赞
 * @param
 * @return
 */
func IsFavoriteRpc(ctx context.Context, user_id int64, video_id_list []int64) (*favorite_rpc.IsFavoriteResp, error) {
	respRpc, err := FavoriteClient.IsFavorite(ctx, &favorite_rpc.IsFavoriteReq{
		UserId:  user_id,
		VideoId: video_id_list,
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
func FavoriteVideoRpc(ctx context.Context, owner_id int64) (*favorite_rpc.FavoriteVideoResp, error) {
	respRpc, err := FavoriteClient.FavoriteVideo(ctx, &favorite_rpc.FavoriteVideoReq{
		UserId: owner_id,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.Code != 0 {
		return nil, fmt.Errorf(respRpc.Msg)
	}

	return respRpc, nil
}
