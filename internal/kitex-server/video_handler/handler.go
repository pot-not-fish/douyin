/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-30 11:46:25
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-22 13:00:56
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\kitex-server\video_handler\handler.go
 */
package video_handler

import (
	"context"
	"douyin/internal/pkg/dal/video_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/kitex_gen/video_rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

type VideoServoceImpl struct{}

func (v *VideoServoceImpl) VideoFeed(ctx context.Context, request *video_rpc.VideoFeedReq) (*video_rpc.VideoFeedResp, error) {
	resp := new(video_rpc.VideoFeedResp)

	klog.CtxDebugf(ctx, "echo called: VideoFeed")

	videoList, nextOffset, err := video_dal.VideoFeed(request.LastOffset, 10)
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

	klog.CtxDebugf(ctx, "echo called: VideoList")

	video_list, err := video_dal.RetrieveUserVideos(request.OwnerId)
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

	klog.CtxDebugf(ctx, "echo called: VideoInfo")

	video_list, err := video_dal.RetrieveVideos(request.VideoId)
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

	video := &video_dal.Video{
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

func (v *VideoServoceImpl) VideoInfoAction(ctx context.Context, request *video_rpc.VideoInfoActionReq) (*video_rpc.VideoInfoActionResp, error) {
	resp := new(video_rpc.VideoInfoActionResp)

	klog.CtxDebugf(ctx, "echo called: VideoInfoAction")

	switch request.ActionType {
	case kitex_client.PubComment:
		if err := video_dal.IncCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DelComment:
		if err := video_dal.DecCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.IncFavorite:
		if err := video_dal.IncFavoriteCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case kitex_client.DecFavorite:
		if err := video_dal.DecFavoriteCount(request.VideoId); err != nil {
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
