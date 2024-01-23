package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/follow_rpc"
	"fmt"
)

func IsFollowRpc(ctx context.Context, user_id_list, to_user_id_list []int64) ([]bool, error) {
	respRpc, err := followClient.IsFollow(ctx, &follow_rpc.IsFollowReq{
		FromUserId: user_id_list,
		ToUserId:   to_user_id_list,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.StatusCode != 0 {
		return nil, fmt.Errorf(respRpc.StatusMsg)
	}

	return respRpc.IsFollow, nil
}
