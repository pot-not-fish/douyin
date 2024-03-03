/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 23:59:20
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 15:10:50
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\main.go
 */
package main

import (
	"douyin/video-rpc/handler"
	"douyin/video-rpc/pkg/dao"
	"douyin/video-rpc/pkg/parse"
	"douyin/video-rpc/video_rpc/videoservice"
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

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8883")
	svr := videoservice.NewServer(
		new(handler.VideoServoceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "video",
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
