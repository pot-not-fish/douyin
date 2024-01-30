package video_handler

import (
	"context"
	"douyin/internal/pkg/kitex_gen/video_rpc"
)

type VideoServoceImpl struct{}

func (v *VideoServoceImpl) VideoList(ctx context.Context, request *video_rpc.VideoListReq) (*video_rpc.VideoListResp, error)

func (v *VideoServoceImpl) VideoInfo(ctx context.Context, request *video_rpc.VideoInfoReq) (*video_rpc.VideoListResp, error)

func (v *VideoServoceImpl) VideoAction(ctx context.Context, request *video_rpc.VideoActionReq) (*video_rpc.VideoActionResp, error)
