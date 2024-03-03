/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 21:57:34
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 23:43:09
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\favorite-mq\pkg\kitex_client\video,go
 */
package kitex_client

import (
	"context"
	"douyin/favorite-mq/pkg/kitex_gen/video_rpc"
	"fmt"
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
