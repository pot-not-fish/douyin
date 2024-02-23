/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-27 10:39:56
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 19:05:15
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\kitex-server\favorite_handler\handler.go
 */
package favorite_handler

import (
	"context"
	"douyin/internal/pkg/dal/favorite_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/kitex_gen/favorite_rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

type FavoriteServiceImpl struct{}

func (f *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite_rpc.FavoriteActionReq) (*favorite_rpc.FavoriteActionResp, error) {
	var err error
	resp := new(favorite_rpc.FavoriteActionResp)

	klog.CtxDebugf(ctx, "echo called: FavoriteAction")

	favorite := favorite_dal.Favorite{
		UserID:  request.UserId,
		VideoID: request.VideoId,
	}

	switch request.ActionType {
	case kitex_client.IncFavorite:
		if err = favorite.CreateFavorite(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DecFavorite:
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

	klog.CtxDebugf(ctx, "echo called: IsFavorite")

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

	klog.CtxDebugf(ctx, "echo called: FavoriteVideo")

	favoriteVideos, err := favorite_dal.RetrieveFavorite(request.UserId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp.VideoId = favoriteVideos
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
