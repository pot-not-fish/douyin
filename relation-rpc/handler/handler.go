/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 19:50:20
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 11:27:36
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-rpc\handler\handler.go
 */
package handler

import (
	"context"
	"douyin/relation-rpc/follow_rpc"
	dao "douyin/relation-rpc/pkg/dao"
)

var (
	FollowList int16 = 6

	FollowerList int16 = 7

	IncFollow int16 = 3

	DecFollow int16 = 4
)

type FollowServiceImpl struct{}

func (f *FollowServiceImpl) RelationAction(ctx context.Context, request *follow_rpc.RelationActionReq) (*follow_rpc.RelationActionResp, error) {
	var err error
	resp := new(follow_rpc.RelationActionResp)

	relation := dao.Relation{
		FollowID:   request.FollowId,
		FollowerID: request.UserId,
	}

	switch request.ActionType {
	case IncFollow:
		if err = relation.CreateRelation(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DecFollow:
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

	for _, v := range request.FollowId {
		is_follow, err := dao.IsFollow(request.UserId, v)
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

	var isFollow []int64
	switch request.ActionType {
	case FollowList:
		isFollow, err = dao.RetrieveFollow(request.UserId)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case FollowerList:
		isFollow, err = dao.RetrieveFollower(request.UserId)
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
