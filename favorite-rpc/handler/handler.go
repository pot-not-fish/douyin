/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 11:49:10
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 14:13:10
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\favorite-rpc\handler\handler.go
 */
package handler

import (
	"context"
	"douyin/favorite-rpc/favorite_rpc"
	"douyin/favorite-rpc/pkg/dao"
)

type FavoriteServiceImpl struct{}

var (
	IncFavorite int16 = 1

	DecFavorite int16 = 2
)

func (f *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite_rpc.FavoriteActionReq) (*favorite_rpc.FavoriteActionResp, error) {
	var err error
	resp := new(favorite_rpc.FavoriteActionResp)

	favorite := dao.Favorite{
		UserID:  request.UserId,
		VideoID: request.VideoId,
	}

	switch request.ActionType {
	case IncFavorite:
		if err = favorite.CreateFavorite(); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DecFavorite:
		if err = favorite.DeleteFavorite(); err != nil {
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

	isFavoriteList, err := dao.IsFavorite(request.UserId, request.VideoId)
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

	favoriteVideos, err := dao.RetrieveFavorite(request.UserId)
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
