/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 21:56:48
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 23:00:10
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\favorite-mq\pkg\kitex_client\init.go
 */
package kitex_client

import (
	"douyin/favorite-mq/pkg/kitex_gen/favorite_rpc/favoriteservice"
	"douyin/favorite-mq/pkg/kitex_gen/user_rpc/userservice"
	"douyin/favorite-mq/pkg/kitex_gen/video_rpc/videoservice"
	"douyin/favorite-mq/pkg/parse"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	UserinfoClient userservice.Client

	FavoriteClient favoriteservice.Client

	VideoClient videoservice.Client
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

	FavoriteClient, err = favoriteservice.NewClient("favorite", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	VideoClient, err = videoservice.NewClient("video", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}
}
