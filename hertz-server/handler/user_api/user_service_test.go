package user_api

import (
	user_api "douyin/hertz-server/model/user_api"
	"douyin/hertz-server/pkg/hook"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/hertz-contrib/gzip"
)

func TestRegister(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"username": "",
		"password": "",
	}
	h.POST("/douyin/user/register", Register)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		fmt.Sprintf(
			"/douyin/user/register?username=%s&password=%s",
			payload["username"], payload["password"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data user_api.RegisterResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestUserinfo(t *testing.T) {
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
	h.GET("/douyin/user", Userinfo)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/user?user_id=%s&token=%s",
			payload["user_id"], payload["token"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
		ut.Header{},
	)

	resp := w.Result()
	var data user_api.UserinfoResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}
