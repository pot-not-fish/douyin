package favorite_handler

import (
	"context"
	"douyin/internal/pkg/dal/favorite_dal"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"
)

type FavoriteServiceImpl struct{}

func (f *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite_rpc.FavoriteActionReq) (*favorite_rpc.FavoriteActionResp, error) {
	var err error
	resp := new(favorite_rpc.FavoriteActionResp)

	favorite := favorite_dal.Favorite{
		UserID:  request.UserId,
		VideoID: request.VideoId,
	}
	switch request.ActionType {
	case 1:
		if err = favorite.CreateFavorite(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 2:
		if err = favorite.CreateFavorite(); err != nil {
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

func (f *FavoriteServiceImpl) IsFavorite(ctx context.Context, request *favorite_rpc.IsFavoriteReq) (*favorite_rpc.IsFavoriteResp, error) {
	resp := new(favorite_rpc.IsFavoriteResp)

	isFavoriteList, err := favorite_dal.IsFavorite(request.UserId, request.VideoId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp.IsFavorite = isFavoriteList
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (f *FavoriteServiceImpl) FavoriteVideo(ctx context.Context, request *favorite_rpc.FavoriteVideoReq) (*favorite_rpc.FavoriteVideoResp, error) {
	resp := new(favorite_rpc.FavoriteVideoResp)

	return resp, nil
}
