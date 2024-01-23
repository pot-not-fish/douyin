package main

import (
	"douyin/internal/kitex-server/follow_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_gen/follow_rpc/followservice"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	user_dal.Init()
	dal.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8886")
	svr := followservice.NewServer(
		new(follow_handler.FollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "isfollow",
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
