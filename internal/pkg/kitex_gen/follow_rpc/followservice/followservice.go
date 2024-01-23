/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-14 12:03:16
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-17 13:57:08
 * @Description: 
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_gen\follow_rpc\followservice\followservice.go
 */
// Code generated by Kitex v0.7.1. DO NOT EDIT.

package followservice

import (
			"context"
				follow_rpc "douyin/internal/pkg/kitex_gen/follow_rpc"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return followServiceServiceInfo
 }

var followServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FollowService"
	handlerType := (*follow_rpc.FollowService)(nil)
	methods := map[string]kitex.MethodInfo{
		"IsFollow":
			kitex.NewMethodInfo(isFollowHandler, newFollowServiceIsFollowArgs, newFollowServiceIsFollowResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "follow_rpc",
		"ServiceFilePath": "..\\..\\idl\\kitex-idl\\follow_rpc.thrift",
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



func isFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*follow_rpc.FollowServiceIsFollowArgs)
	realResult := result.(*follow_rpc.FollowServiceIsFollowResult)
	success, err := handler.(follow_rpc.FollowService).IsFollow(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newFollowServiceIsFollowArgs() interface{} {
	return follow_rpc.NewFollowServiceIsFollowArgs()
}

func newFollowServiceIsFollowResult() interface{} {
	return follow_rpc.NewFollowServiceIsFollowResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) IsFollow(ctx context.Context , request *follow_rpc.IsFollowReq) (r *follow_rpc.IsFollowResp, err error) {
	var _args follow_rpc.FollowServiceIsFollowArgs
	_args.Request = request
	var _result follow_rpc.FollowServiceIsFollowResult
	if err = p.c.Call(ctx, "IsFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

