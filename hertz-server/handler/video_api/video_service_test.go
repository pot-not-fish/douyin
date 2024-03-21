package video_api

import (
	"douyin/hertz-server/model/video_api"
	"douyin/hertz-server/pkg/hook"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/hertz-contrib/gzip"
)

func TestFeed(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"latest_time": "",
		"token":       "",
	}
	h.GET("/douyin/feed", Feed)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/feed?token=%s&latest_time=%s",
			payload["token"], payload["latest_time"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data video_api.FeedResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestList(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"user_id": "",
		"token":   "",
	}
	h.GET("/douyin/feed", List)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/feed?user_id=%s&token=%s",
			payload["user_id"], payload["token"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data video_api.ListResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

// TODO: 这个文件的接口的单测没有写出来
func TestPublish(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	h.POST("/douyin/publish", Publish)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		"/douyin/publish",
		&ut.Body{},
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data PublishResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}
