package follow_handler

import (
	"context"
	"douyin/internal/pkg/dal/relation_dal"
	"douyin/internal/pkg/kitex_gen/follow_rpc"
)

type FollowServiceImpl struct{}

func (f *FollowServiceImpl) RelationAction(ctx context.Context, request *follow_rpc.RelationActionReq) (*follow_rpc.RelationActionResp, error) {
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

func (f *FollowServiceImpl) IsFollow(ctx context.Context, request *follow_rpc.IsFollowReq) (*follow_rpc.IsFollowResp, error) {
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

func (f *FollowServiceImpl) RelationList(ctx context.Context, request *follow_rpc.RelationListReq) (*follow_rpc.RelationListResp, error) {
	var err error
	resp := new(follow_rpc.RelationListResp)

	relation_info := new(relation_dal.RelationInfo)
	switch request.ActionType {
	case 1:
		relation_info, err = relation_dal.RetrieveFollow(request.UserId, request.OwnerId)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 2:
		relation_info, err = relation_dal.RetrieveFollower(request.UserId, request.OwnerId)
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

	resp.IsFollow = relation_info.IsFollowList
	resp.UserId = relation_info.RelationList
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
