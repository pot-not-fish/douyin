// Code generated by hertz generator. DO NOT EDIT.

package video_api

import (
	video_api "douyin/hertz-server/handler/video_api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), video_api.Feed)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_list0Mw(), video_api.List)...)
			}
		}
	}
}