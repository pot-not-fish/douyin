/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-13 19:38:25
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-12 19:59:00
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\hertz-server\pkg\kitex_client\init.go
 */
package kitex_client

import (
	"douyin/hertz-server/pkg/kitex_gen/comment_rpc/commentservice"
	"douyin/hertz-server/pkg/kitex_gen/favorite_rpc/favoriteservice"
	"douyin/hertz-server/pkg/kitex_gen/follow_rpc/followservice"
	"douyin/hertz-server/pkg/kitex_gen/user_rpc/userservice"
	"douyin/hertz-server/pkg/kitex_gen/video_rpc/videoservice"
	"douyin/hertz-server/pkg/parse"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	userinfoClient userservice.Client

	favoriteClient favoriteservice.Client

	followClient followservice.Client

	videoClient videoservice.Client

	commentClient commentservice.Client

	once *sync.Once
)

func Init() {
	var err error

	once.Do(func() {
		etcdClient, err = etcd.NewEtcdResolver([]string{parse.ConfigStructure.Etcd.Host})
		if err != nil {
			panic(err)
		}

		userinfoClient, err = userservice.NewClient("user", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}

		favoriteClient, err = favoriteservice.NewClient("favorite", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}

		followClient, err = followservice.NewClient("follow", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}

		videoClient, err = videoservice.NewClient("video", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}

		commentClient, err = commentservice.NewClient("comment", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}
	})
}
