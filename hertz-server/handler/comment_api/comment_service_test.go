package comment_api

import (
	"douyin/hertz-server/model/comment_api"
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

func TestCommentAction(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"token":        "",
		"video_id":     "",
		"action_type":  "",
		"comment_text": "",
		"comment_id":   "",
	}
	h.POST("/douyin/comment/action", mw.JwtMiddleware.MiddlewareFunc(), CommentAction)
	w := ut.PerformRequest(
		h.Engine,
		"POST",
		fmt.Sprintf(
			"/douyin/comment/action?token=%s&video_id=%s&action_type=%s&comment_text=%s&comment_id=%s",
			payload["token"], payload["video_id"], payload["action_type"], payload["comment_text"], payload["comment_id"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data comment_api.CommentActionResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}

func TestCommentList(t *testing.T) {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithMaxRequestBodySize(50*1024*1024),
	)
	hook.StartHook(h)
	hook.ShutdownHook(h)
	h.Use(gzip.Gzip(gzip.BestCompression))

	var payload = map[string]string{
		"token":    "",
		"video_id": "",
	}
	h.GET("/douyin/comment/list", mw.JwtMiddleware.MiddlewareFunc(), CommentList)
	w := ut.PerformRequest(
		h.Engine,
		"GET",
		fmt.Sprintf(
			"/douyin/comment/list?token=%s&video_id=%s",
			payload["token"], payload["video_id"],
		),
		nil,
		ut.Header{Key: "Connection", Value: "close"},
	)

	resp := w.Result()
	var data comment_api.CommentListResp
	sonic.Unmarshal(resp.Body(), &data)
	assert.DeepEqual(t, 200, resp.StatusCode)
	assert.DeepEqual(t, 0, data.StatusCode)
}
