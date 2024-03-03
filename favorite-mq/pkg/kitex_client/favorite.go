package kitex_client

import (
	"context"
	"douyin/favorite-mq/pkg/kitex_gen/favorite_rpc"
	"fmt"
)

var (
	IncFavorite int16 = 1

	DecFavorite int16 = 2
)

/**
 * @function
 * @description 点赞视频操作
 * @param action_type 操作码 IncFavorite-点赞，DecFavorite-取消点赞
 * @return
 */
func FavoriteActionRpc(ctx context.Context, action_type int16, user_id, video_id int64) error {
	respRpc, err := FavoriteClient.FavoriteAction(ctx, &favorite_rpc.FavoriteActionReq{
		ActionType: action_type,
		UserId:     user_id,
		VideoId:    video_id,
	})
	if err != nil {
		return err
	}

	if respRpc.Code != 0 {
		return fmt.Errorf(respRpc.Msg)
	}

	return nil
}
