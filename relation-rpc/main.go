/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 23:43:13
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 12:15:59
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-rpc\main.go
 */
package main

import (
	"douyin/relation-rpc/follow_rpc/followservice"
	"douyin/relation-rpc/handler"
	dao "douyin/relation-rpc/pkg/dao"
	"douyin/relation-rpc/pkg/parse"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// parse.Init("./local.yaml")
	parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{parse.ConfigStructure.Etcd.Host})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8885")
	svr := followservice.NewServer(
		new(handler.FollowServiceImpl),
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
