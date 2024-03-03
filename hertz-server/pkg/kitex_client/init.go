/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-13 19:38:25
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 11:09:53
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

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	UserinfoClient userservice.Client

	FavoriteClient favoriteservice.Client

	FollowClient followservice.Client

	VideoClient videoservice.Client

	CommentClient commentservice.Client
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

	FollowClient, err = followservice.NewClient("follow", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	VideoClient, err = videoservice.NewClient("video", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}

	CommentClient, err = commentservice.NewClient("comment", client.WithResolver(etcdClient))
	if err != nil {
		panic(err)
	}
}
