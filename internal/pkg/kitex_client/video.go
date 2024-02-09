/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-02 13:46:12
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-08 15:06:26
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_client\video.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/video_rpc"
	"fmt"
)

/**
 * @function
 * @description 查看视频流
 * @param
 * @return
 */
func VideoFeedRpc(ctx context.Context, last_offset int64) (*video_rpc.VideoFeedResp, error) {
	respRpc, err := VideoClient.VideoFeed(ctx, &video_rpc.VideoFeedReq{
		LastOffset: last_offset,
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
 * @description 查看某个人发布的视频列表
 * @param
 * @return
 */
func VideoListRpc(ctx context.Context, owner_id int64) (*video_rpc.VideoListResp, error) {
	respRpc, err := VideoClient.VideoList(ctx, &video_rpc.VideoListReq{
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

/**
 * @function
 * @description 查看一组视频信息
 * @param
 * @return
 */
func VideoInfoRpc(ctx context.Context, video_id []int64) (*video_rpc.VideoListResp, error) {
	respRpc, err := VideoClient.VideoInfo(ctx, &video_rpc.VideoInfoReq{
		VideoId: video_id,
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
 * @description 发布视频操作
 * @param
 * @return
 */
func VideoActionRpc(ctx context.Context, user_id int64, title, cover_url, play_url string) error {
	respRpc, err := VideoClient.VideoAction(ctx, &video_rpc.VideoActionReq{
		UserId:   user_id,
		Title:    title,
		CoverUrl: cover_url,
		PlayUrl:  play_url,
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
	IncVideoComment int16 = 1

	DecVideoComment int16 = 2

	IncVideoFavorite int16 = 3

	DecVideoFavorite int16 = 4
)

/**
 * @function
 * @description 评论自增、自减，点赞自增、自减操作
 * @param action_type 操作码 IncVideoComment-自增评论 DecVideoComment-自减评论 IncVideoFavorite-自增点赞量 DecVideoFavorite-自减点赞量
 * @return
 */
func VideoInfoActionRpc(ctx context.Context, action_type int16, video_id int64) error {
	respRpc, err := VideoClient.VideoInfoAction(ctx, &video_rpc.VideoInfoActionReq{
		ActionType: action_type,
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
