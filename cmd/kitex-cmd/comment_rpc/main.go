/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-30 14:36:56
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-22 13:08:24
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\cmd\kitex-cmd\comment_rpc\main.go
 */
package main

import (
	"douyin/internal/kitex-server/comment_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/comment_dal"
	"douyin/internal/pkg/kitex_gen/comment_rpc/commentservice"
	"douyin/internal/pkg/parse"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	parse.Init("../../../deployment/config/config.yaml")

	comment_dal.Init()

	dal.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8887")
	svr := commentservice.NewServer(
		new(comment_handler.CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "comment",
		}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithReadWriteTimeout(time.Hour), // 如果client多次发送数据，间隔较长，需要延长服务端的超时时间
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
