package follow_handler

import (
	"context"
	"douyin/internal/pkg/dal/relation_dal"
	"douyin/internal/pkg/kitex_gen/follow_rpc"
)

type FollowServiceImpl struct{}

func (r *FollowServiceImpl) RelationAction(ctx context.Context, request *follow_rpc.RelationActionReq) (*follow_rpc.RelationActionResp, error) {
	var err error
	resp := new(follow_rpc.RelationActionResp)

	relation := relation_dal.Relation{
		FollowID:   request.FollowId,
		FollowerID: request.UserId,
	}

	switch request.ActionType {
	case 1:
		if err = relation.CreateRelation(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 2:
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

func (r *FollowServiceImpl) IsFollow(ctx context.Context, request *follow_rpc.IsFollowReq) (*follow_rpc.IsFollowResp, error) {
	resp := new(follow_rpc.IsFollowResp)

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

func (r *FollowServiceImpl) FollowList(ctx context.Context, request *follow_rpc.FollowListReq) (*follow_rpc.FollowListResp, error) {
	resp := new(follow_rpc.FollowListResp)

	relation_info, err := relation_dal.RetrieveFollow(request.UserId, request.OwnerId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp = &follow_rpc.FollowListResp{
		Code:     0,
		Msg:      "ok",
		IsFollow: relation_info.IsFollowList,
		UserId:   relation_info.RelationList,
	}
	return resp, nil
}

func (r *FollowServiceImpl) FollowerList(ctx context.Context, request *follow_rpc.FollowerListReq) (*follow_rpc.FollowerListResp, error) {
	resp := new(follow_rpc.FollowerListResp)

	relation_info, err := relation_dal.RetrieveFollower(request.UserId, request.OwnerId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp = &follow_rpc.FollowerListResp{
		Code:     0,
		Msg:      "ok",
		IsFollow: relation_info.IsFollowList,
		UserId:   relation_info.RelationList,
	}
	return resp, nil
}
