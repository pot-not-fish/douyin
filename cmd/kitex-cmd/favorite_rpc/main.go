/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-27 10:42:13
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-18 11:06:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\cmd\kitex-cmd\favorite_rpc\main.go
 */
package main

import (
	"douyin/internal/kitex-server/favorite_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/favorite_dal"
	"douyin/internal/pkg/kitex_gen/favorite_rpc/favoriteservice"
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

	favorite_dal.Init()

	dal.InitRedis()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8886")
	svr := favoriteservice.NewServer(
		new(favorite_handler.FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "favorite",
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
