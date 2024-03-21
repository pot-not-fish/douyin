/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 20:00:46
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-21 14:48:34
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\comment-rpc\main.go
 */
package main

import (
	"douyin/comment-rpc/comment_rpc/commentservice"
	"douyin/comment-rpc/handler"
	"douyin/comment-rpc/pkg/dao"
	"douyin/comment-rpc/pkg/parse"
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

	r, err := etcd.NewEtcdRegistry([]string{parse.ConfigStructure.Etcd.Host})
	if err != nil {
		log.Println(err.Error())
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8887")
	svr := commentservice.NewServer(
		new(handler.CommentServiceImpl),
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
