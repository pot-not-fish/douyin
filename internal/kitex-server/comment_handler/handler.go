package comment_handler

import (
	"context"
	"douyin/internal/pkg/dal/comment_dal"
	"douyin/internal/pkg/kitex_gen/comment_rpc"
	"fmt"
)

type CommentServiceImpl struct{}

func (c *CommentServiceImpl) CommentAction(ctx context.Context, request *comment_rpc.CommentActionReq) (*comment_rpc.CommentActionResp, error) {
	var err error
	resp := new(comment_rpc.CommentActionResp)

	comment := comment_dal.Comment{
		UserID:  request.UserId,
		VideoID: request.VideoId,
	}

	switch request.ActionType {
	case 1:
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
	case 2:
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
	resp := new(comment_rpc.CommentListResp)

	comment_list, err := comment_dal.RetrieveComment(request.VideoId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	for _, v := range comment_list {
		resp.Comment = append(resp.Comment, &comment_rpc.Comment{
			Id:         int64(v.ID),
			Content:    v.Content,
			CreateDate: fmt.Sprintf("%d-%d-%d %d:%d", v.CreatedAt.Year(), v.CreatedAt.Month(), v.CreatedAt.Day(), v.CreatedAt.Hour(), v.CreatedAt.Minute()),
		})
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
