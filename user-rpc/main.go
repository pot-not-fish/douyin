/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 23:33:44
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 15:10:34
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\user-rpc\main.go
 */
package main

import (
	"douyin/user-rpc/handler"
	dao "douyin/user-rpc/pkg/dao"
	"douyin/user-rpc/pkg/parse"
	"douyin/user-rpc/user_rpc/userservice"
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

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8884")
	svr := userservice.NewServer(
		new(handler.UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user",
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
