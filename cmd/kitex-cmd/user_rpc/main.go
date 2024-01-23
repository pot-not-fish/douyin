/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-13 10:36:48
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2023-12-25 10:53:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\cmd\kitex-cmd\user_rpc\main.go
 */
package main

import (
	"douyin/internal/kitex-server/user_handler"
	"douyin/internal/pkg/dal"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_gen/user_rpc/userservice"
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

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8885")
	svr := userservice.NewServer(
		new(user_handler.UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "userinfo",
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
