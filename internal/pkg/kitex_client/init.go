/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-13 19:38:25
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-14 20:56:46
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_client\init.go
 */
package kitex_client

import (
	"douyin/internal/pkg/kitex_gen/favorite_rpc/favoriteservice"
	"douyin/internal/pkg/kitex_gen/follow_rpc/followservice"
	"douyin/internal/pkg/kitex_gen/user_rpc/userservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	userinfoClient userservice.Client

	favoriteClient favoriteservice.Client

	followClient followservice.Client
)

func Init() {
	var err error
	etcdClient, err = etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	userinfoClient, err = userservice.NewClient("userinfo", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	favoriteClient, err = favoriteservice.NewClient("isfavorite", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	followClient, err = followservice.NewClient("isfollow", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}
}
