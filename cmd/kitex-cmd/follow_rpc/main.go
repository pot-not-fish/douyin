/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-14 12:04:14
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-13 17:31:55
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\cmd\kitex-cmd\follow_rpc\main.go
 */
package main

import (
	"douyin/internal/kitex-server/follow_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/relation_dal"
	"douyin/internal/pkg/kitex_gen/follow_rpc/followservice"
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
	relation_dal.Init()
	dal.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8885")
	svr := followservice.NewServer(
		new(follow_handler.FollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "follow",
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
