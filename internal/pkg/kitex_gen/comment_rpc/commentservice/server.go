// Code generated by Kitex v0.7.1. DO NOT EDIT.
package commentservice

import (
	comment_rpc "douyin/internal/pkg/kitex_gen/comment_rpc"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler comment_rpc.CommentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
