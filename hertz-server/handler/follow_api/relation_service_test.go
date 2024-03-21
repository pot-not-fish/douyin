package follow_api

import (
	"douyin/hertz-server/model/follow_api"
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

func TestRelationAction(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"token":       "",
		"to_user_id":  "",
		"action_type": "",
	}
	h.POST("/douyin/relation/action", mw.JwtMiddleware.MiddlewareFunc(), RelationAction)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		fmt.Sprintf(
			"/douyin/relation/action?token=%s&to_user_id=%s&action_type=%s",
			payload["token"], payload["to_user_id"], payload["action_type"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data follow_api.RelationActionResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestRelationFollow(t *testing.T) {
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
	h.GET("/douyin/relatioin/follow/list", RelationFollow)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/relatioin/follow/list?token=%s&user_id=%s",
			payload["token"], payload["user_id"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data follow_api.FollowListResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestRelationFollower(t *testing.T) {
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
	h.GET("/douyin/relation/follower/list", RelationFollower)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/relation/follower/list?token=%s&user_id=%s",
			payload["token"], payload["user_id"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data follow_api.FollowerListResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}
