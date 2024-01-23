package favorite_handler

import (
	"context"
	"douyin/internal/pkg/dal/video_dal"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"
)

type FavoriteServiceImpl struct{}

func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, request *favorite_rpc.IsFavoriteReq) (resp *favorite_rpc.IsFavoriteResp, err error) {
	isfavorite_list := make([]bool, 0, len(request.VideoId))

	for k, v := range request.UserId {
		favorite, err := video_dal.RetrieveFavorite(v, request.VideoId[k])
		if err != nil {
			return &favorite_rpc.IsFavoriteResp{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}, nil
		}

		isfavorite_list = append(isfavorite_list, favorite)
	}

	return &favorite_rpc.IsFavoriteResp{
		StatusCode: 0,
		StatusMsg:  "OK",
		IsFavorite: isfavorite_list,
	}, nil
}
