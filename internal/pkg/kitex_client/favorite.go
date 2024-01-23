package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"
	"fmt"
)

func IsFavoriteRpc(ctx context.Context, userid []int64, videoid []int64) ([]bool, error) {
	respRpc, err := favoriteClient.IsFavorite(ctx, &favorite_rpc.IsFavoriteReq{
		UserId:  userid,
		VideoId: videoid,
	})
	if err != nil {
		return nil, err
	}

	if respRpc.StatusCode != 0 {
		return nil, fmt.Errorf(respRpc.StatusMsg)
	}

	return respRpc.IsFavorite, nil
}
