package favorite_api

import (
	"douyin/hertz-server/model/favorite_api"
	"douyin/hertz-server/pkg/hook"
	"douyin/hertz-server/pkg/mw"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/hertz-contrib/gzip"
)

func TestFavoriteAction(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"token":       "",
		"video_id":    "",
		"action_type": "",
	}
	h.POST("/douyin/favorite/action", mw.JwtMiddleware.MiddlewareFunc(), FavoriteAction)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		fmt.Sprintf("/douyin/favorite/action?token=%s&video_id=%s&action_type=%s", payload["token"], payload["video_id"], payload["action_type"]),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data favorite_api.FavoriteActionResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestFavoriteList(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"token":   "",
		"user_id": "",
	}
	h.POST("/douyin/favorite/list", FavoriteList)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		fmt.Sprintf("/douyin/favorite/list?token=%s&user_id=%s", payload["token"], payload["user_id"]),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data favorite_api.FavoriteListResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}
