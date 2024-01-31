package video_handler

import (
	"context"
	"douyin/internal/pkg/kitex_gen/video_rpc"
)

type VideoServoceImpl struct{}

func (v *VideoServoceImpl) VideoList(ctx context.Context, request *video_rpc.VideoListReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

	return resp, nil
}

func (v *VideoServoceImpl) VideoInfo(ctx context.Context, request *video_rpc.VideoInfoReq) (*video_rpc.VideoListResp, error) {
	resp := new(video_rpc.VideoListResp)

	return resp, nil
}

func (v *VideoServoceImpl) VideoAction(ctx context.Context, request *video_rpc.VideoActionReq) (*video_rpc.VideoActionResp, error) {
	resp := new(video_rpc.VideoActionResp)

	return resp, nil
}

func (v *VideoServoceImpl) VideoInfoAction(ctx context.Context, request *video_rpc.VideoInfoActionReq) (*video_rpc.VideoInfoActionResp, error) {
	resp := new(video_rpc.VideoInfoActionResp)

	return resp, nil
}
