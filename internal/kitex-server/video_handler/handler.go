/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-30 11:46:25
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-02 13:43:03
 * @Description:
 * @FilePath: \douyin\internal\kitex-server\video_handler\handler.go
 */
package video_handler

import (
	"context"
	"douyin/internal/pkg/dal/video_dal"
	"douyin/internal/pkg/kitex_gen/video_rpc"
)

type VideoServoceImpl struct{}

func (v *VideoServoceImpl) VideoFeed(ctx context.Context, request *video_rpc.VideoFeedReq) (*video_rpc.VideoFeedResp, error) {
	resp := new(video_rpc.VideoFeedResp)

	video_list, next_offset, err := video_dal.VideoFeed(request.LastOffset, 10)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, nil
	}

	resp.NextOffset = next_offset
	for _, v := range video_list {
		resp.Videos = append(resp.Videos, &video_rpc.Video{
			Id:            int64(v.ID),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
		})
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoList(ctx context.Context, request *video_rpc.VideoListReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

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
		})
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoInfo(ctx context.Context, request *video_rpc.VideoInfoReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

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
		})
	}
	resp.Code = 0
	resp.Msg = "ok"
	return resp, nil
}

func (v *VideoServoceImpl) VideoAction(ctx context.Context, request *video_rpc.VideoActionReq) (*video_rpc.VideoActionResp, error) {
	resp := new(video_rpc.VideoActionResp)

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

	switch request.ActionType {
	case 1:
		if err := video_dal.IncCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 2:
		if err := video_dal.DecCommentCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 3:
		if err := video_dal.IncFavoriteCount(request.VideoId); err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
			return resp, nil
		}
	case 4:
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
