/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-02 13:46:28
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-02 18:38:06
 * @Description:
 * @FilePath: \douyin\internal\pkg\kitex_client\comment.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/comment_rpc"
	"fmt"
)

/**
 * @function
 * @description 发布评论操作
 * @param action_type 操作码 1-评论，2-删除评论
 * @return
 */
func CommentActionRpc(ctx context.Context, action_type int16, user_id, video_id int64, content *string) (*comment_rpc.CommentActionResp, error) {
	respRpc, err := CommentClient.CommentAction(ctx, &comment_rpc.CommentActionReq{
		ActionType:  action_type,
		UserId:      user_id,
		VideoId:     video_id,
		CommentText: content,
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
 * @description 查看视频的评论
 * @param
 * @return
 */
func CommentListRpc(ctx context.Context, video_id int64) (*comment_rpc.CommentListResp, error) {
	respRpc, err := CommentClient.CommentList(ctx, &comment_rpc.CommentListReq{
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
