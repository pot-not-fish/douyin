package follow_handler

import (
	"context"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_gen/follow_rpc"
)

type FollowServiceImpl struct{}

func (s *FollowServiceImpl) IsFollow(ctx context.Context, request *follow_rpc.IsFollowReq) (resp *follow_rpc.IsFollowResp, err error) {
	isfollow_list := make([]bool, 0, len(request.ToUserId))

	for k, v := range request.FromUserId {
		isfollow := user_dal.IsFollow(v, request.ToUserId[k])
		isfollow_list = append(isfollow_list, isfollow)
	}

	return &follow_rpc.IsFollowResp{
		StatusCode: 0,
		StatusMsg:  "OK",
		IsFollow:   isfollow_list,
	}, nil
}
