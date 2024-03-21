/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 23:33:44
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-21 15:03:39
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\user-rpc\main.go
 */
package main

import (
	"douyin/user-rpc/handler"
	"douyin/user-rpc/pkg/dao"
	"douyin/user-rpc/pkg/parse"
	"douyin/user-rpc/user_rpc/userservice"
	"log"
	"net"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	var err error
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)
	klog.SetLevel(klog.LevelDebug)

	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

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
