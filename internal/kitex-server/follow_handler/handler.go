/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-14 12:05:03
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-22 13:01:42
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\kitex-server\follow_handler\handler.go
 */
package follow_handler

import (
	"context"
	"douyin/internal/pkg/dal/relation_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/kitex_gen/follow_rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

type FollowServiceImpl struct{}

func (f *FollowServiceImpl) RelationAction(ctx context.Context, request *follow_rpc.RelationActionReq) (*follow_rpc.RelationActionResp, error) {
	var err error
	resp := new(follow_rpc.RelationActionResp)

	klog.CtxDebugf(ctx, "echo called: RelationAction")

	relation := relation_dal.Relation{
		FollowID:   request.FollowId,
		FollowerID: request.UserId,
	}

	switch request.ActionType {
	case kitex_client.IncFollow:
		if err = relation.CreateRelation(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DecFollow:
		if err = relation.DeleteRelation(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "invalid action type"
		return resp, nil
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (f *FollowServiceImpl) IsFollow(ctx context.Context, request *follow_rpc.IsFollowReq) (*follow_rpc.IsFollowResp, error) {
	resp := new(follow_rpc.IsFollowResp)

	klog.CtxDebugf(ctx, "echo called: IsFollow")

	for _, v := range request.FollowId {
		is_follow, err := relation_dal.IsFollow(request.UserId, v)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
		resp.IsFollow = append(resp.IsFollow, is_follow)
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (f *FollowServiceImpl) RelationList(ctx context.Context, request *follow_rpc.RelationListReq) (*follow_rpc.RelationListResp, error) {
	var err error
	resp := new(follow_rpc.RelationListResp)

	klog.CtxDebugf(ctx, "echo called: RelationList")

	var isFollow []int64
	switch request.ActionType {
	case kitex_client.FollowList:
		isFollow, err = relation_dal.RetrieveFollow(request.UserId)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.FollowerList:
		isFollow, err = relation_dal.RetrieveFollower(request.UserId)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "invalid action type"
		return resp, nil
	}

	resp.UserId = isFollow
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
