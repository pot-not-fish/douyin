// Code generated by Kitex v0.7.1. DO NOT EDIT.

package videoservice

import (
			"context"
				video_rpc "douyin/internal/pkg/kitex_gen/video_rpc"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
 }

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video_rpc.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"VideoList":
			kitex.NewMethodInfo(videoListHandler, newVideoServiceVideoListArgs, newVideoServiceVideoListResult, false),
		"VideoInfo":
			kitex.NewMethodInfo(videoInfoHandler, newVideoServiceVideoInfoArgs, newVideoServiceVideoInfoResult, false),
		"VideoAction":
			kitex.NewMethodInfo(videoActionHandler, newVideoServiceVideoActionArgs, newVideoServiceVideoActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "video_rpc",
		"ServiceFilePath": "..\\..\\idl\\kitex-idl\\video_rpc.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName: 	 serviceName,
		HandlerType: 	 handlerType,
		Methods:     	 methods,
		PayloadCodec:  	 kitex.Thrift,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}



func videoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*video_rpc.VideoServiceVideoListArgs)
	realResult := result.(*video_rpc.VideoServiceVideoListResult)
	success, err := handler.(video_rpc.VideoService).VideoList(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newVideoServiceVideoListArgs() interface{} {
	return video_rpc.NewVideoServiceVideoListArgs()
}

func newVideoServiceVideoListResult() interface{} {
	return video_rpc.NewVideoServiceVideoListResult()
}


func videoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*video_rpc.VideoServiceVideoInfoArgs)
	realResult := result.(*video_rpc.VideoServiceVideoInfoResult)
	success, err := handler.(video_rpc.VideoService).VideoInfo(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newVideoServiceVideoInfoArgs() interface{} {
	return video_rpc.NewVideoServiceVideoInfoArgs()
}

func newVideoServiceVideoInfoResult() interface{} {
	return video_rpc.NewVideoServiceVideoInfoResult()
}


func videoActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*video_rpc.VideoServiceVideoActionArgs)
	realResult := result.(*video_rpc.VideoServiceVideoActionResult)
	success, err := handler.(video_rpc.VideoService).VideoAction(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newVideoServiceVideoActionArgs() interface{} {
	return video_rpc.NewVideoServiceVideoActionArgs()
}

func newVideoServiceVideoActionResult() interface{} {
	return video_rpc.NewVideoServiceVideoActionResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) VideoList(ctx context.Context , request *video_rpc.VideoListReq) (r *video_rpc.VideoListResp, err error) {
	var _args video_rpc.VideoServiceVideoListArgs
	_args.Request = request
	var _result video_rpc.VideoServiceVideoListResult
	if err = p.c.Call(ctx, "VideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoInfo(ctx context.Context , request *video_rpc.VideoInfoReq) (r *video_rpc.VideoListResp, err error) {
	var _args video_rpc.VideoServiceVideoInfoArgs
	_args.Request = request
	var _result video_rpc.VideoServiceVideoInfoResult
	if err = p.c.Call(ctx, "VideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoAction(ctx context.Context , request *video_rpc.VideoActionReq) (r *video_rpc.VideoActionResp, err error) {
	var _args video_rpc.VideoServiceVideoActionArgs
	_args.Request = request
	var _result video_rpc.VideoServiceVideoActionResult
	if err = p.c.Call(ctx, "VideoAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
