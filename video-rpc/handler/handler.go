/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 23:59:48
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 11:35:52
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\handler\handler.go
 */
package handler

import (
	"context"
	"douyin/video-rpc/pkg/dao"
	"douyin/video-rpc/video_rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

type VideoServoceImpl struct{}

func (v *VideoServoceImpl) VideoFeed(ctx context.Context, request *video_rpc.VideoFeedReq) (*video_rpc.VideoFeedResp, error) {
	resp := new(video_rpc.VideoFeedResp)

	videoList, nextOffset, err := dao.VideoFeed(request.LastOffset, 10)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	for _, v := range videoList {
		resp.Videos = append(resp.Videos, &video_rpc.Video{
			Id:            int64(v.ID),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			UserId:        v.UserID,
		})
	}
	resp.NextOffset = nextOffset
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoList(ctx context.Context, request *video_rpc.VideoListReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

	video_list, err := dao.RetrieveUserVideos(request.OwnerId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	for _, v := range video_list {
		resp.Videos = append(resp.Videos, &video_rpc.Video{
			Id:            int64(v.ID),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			UserId:        v.UserID,
		})
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoInfo(ctx context.Context, request *video_rpc.VideoInfoReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

	video_list, err := dao.RetrieveVideos(request.VideoId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	for _, v := range video_list {
		resp.Videos = append(resp.Videos, &video_rpc.Video{
			Id:            int64(v.ID),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			UserId:        v.UserID,
		})
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoAction(ctx context.Context, request *video_rpc.VideoActionReq) (*video_rpc.VideoActionResp, error) {
	resp := new(video_rpc.VideoActionResp)

	klog.CtxDebugf(ctx, "echo called: VideoAction")

	video := &dao.Video{
		UserID:   request.UserId,
		PlayUrl:  request.PlayUrl,
		CoverUrl: request.CoverUrl,
		Title:    request.Title,
	}

	if err := video.CreateVideo(); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

var (
	IncFavorite int16 = 1

	DecFavorite int16 = 2

	PubComment int16 = 8

	DelComment int16 = 9
)

func (v *VideoServoceImpl) VideoInfoAction(ctx context.Context, request *video_rpc.VideoInfoActionReq) (*video_rpc.VideoInfoActionResp, error) {
	resp := new(video_rpc.VideoInfoActionResp)

	klog.CtxDebugf(ctx, "echo called: VideoInfoAction")

	switch request.ActionType {
	case PubComment:
		if err := dao.IncCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DelComment:
		if err := dao.DecCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case IncFavorite:
		if err := dao.IncFavoriteCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case DecFavorite:
		if err := dao.DecFavoriteCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	default:
		resp.Code = 1
		resp.Msg = "invalid action type"
	}

	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}
