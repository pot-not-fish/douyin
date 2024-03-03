/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 21:41:54
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 23:52:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-mq\pkg\kitex_client\init.go
 */
package kitex_client

import (
	"douyin/relation-mq/pkg/kitex_gen/follow_rpc/followservice"
	"douyin/relation-mq/pkg/kitex_gen/user_rpc/userservice"
	"douyin/relation-mq/pkg/parse"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	UserinfoClient userservice.Client

	FollowClient followservice.Client
)

func Init() {
	var err error
	etcdClient, err = etcd.NewEtcdResolver([]string{parse.ConfigStructure.Etcd.Host})
	if err != nil {
		panic(err)
	}

	UserinfoClient, err = userservice.NewClient("user", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	FollowClient, err = followservice.NewClient("follow", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}
}
