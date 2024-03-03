/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 19:59:08
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-28 20:04:23
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\comment-rpc\handler\handler.go
 */
package handler

import (
	"context"
	"douyin/comment-rpc/comment_rpc"
	"douyin/comment-rpc/pkg/dao"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CommentServiceImpl struct{}

var (
	PubComment int16 = 8

	DelComment int16 = 9
)

func (c *CommentServiceImpl) CommentAction(ctx context.Context, request *comment_rpc.CommentActionReq) (*comment_rpc.CommentActionResp, error) {
	var err error
	resp := new(comment_rpc.CommentActionResp)

	klog.CtxDebugf(ctx, "echo called: CommentAction")

	comment := dao.Comment{
		UserID:  request.UserId,
		VideoID: request.VideoId,
	}

	switch request.ActionType {
	case PubComment:
		if request.CommentText == nil || *request.CommentText == "" {
			resp.Code = 1
			resp.Msg = "Invalid comment text"
			return resp, nil
		}

		comment.Content = *request.CommentText
		if err = comment.CreateComment(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}

		create_date := fmt.Sprintf("%d-%d-%d %d:%d", comment.CreatedAt.Year(), comment.CreatedAt.Month(), comment.CreatedAt.Day(), comment.CreatedAt.Hour(), comment.CreatedAt.Minute())
		resp.Comment = &comment_rpc.Comment{
			Id:         int64(comment.ID),
			Content:    *request.CommentText,
			CreateDate: create_date,
			UserId:     request.UserId,
		}
	case DelComment:
		if err = comment.DeleteComment(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "Invalid action type"
		return resp, nil
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (c *CommentServiceImpl) CommentList(ctx context.Context, request *comment_rpc.CommentListReq) (*comment_rpc.CommentListResp, error) {
	var err error
	resp := new(comment_rpc.CommentListResp)

	klog.CtxDebugf(ctx, "echo called: CommentList")

	comment_list, err := dao.RetrieveComment(request.VideoId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	if len(comment_list) == 0 {
		resp.Code = 0
		resp.Msg = "ok"
		return resp, nil
	}

	for _, v := range comment_list {
		resp.Comment = append(resp.Comment, &comment_rpc.Comment{
			Id:         int64(v.ID),
			Content:    v.Content,
			CreateDate: fmt.Sprintf("%d-%d-%d %d:%d", v.CreatedAt.Year(), v.CreatedAt.Month(), v.CreatedAt.Day(), v.CreatedAt.Hour(), v.CreatedAt.Minute()),
			UserId:     v.UserID,
		})
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
