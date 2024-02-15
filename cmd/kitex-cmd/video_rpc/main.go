package main

import (
	"douyin/internal/kitex-server/video_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/video_dal"
	"douyin/internal/pkg/kitex_gen/video_rpc/videoservice"
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
	video_dal.Init()
	dal.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8883")
	svr := videoservice.NewServer(
		new(video_handler.VideoServoceImpl),
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
