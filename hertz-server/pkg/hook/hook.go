/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-13 10:47:28
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-18 19:31:48
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\hertz-server\pkg\hook\hook.go
 */
package hook

import (
	"context"
	"douyin/hertz-server/pkg/kitex_client"
	"douyin/hertz-server/pkg/mq"
	"douyin/hertz-server/pkg/mw"
	"douyin/hertz-server/pkg/parse"
	"sync"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	startOnce    *sync.Once
	shutdownOnce *sync.Once
)

func StartHook(s *server.Hertz) {
	startOnce.Do(func() {
		s.OnRun = append(s.OnRun, func(ctx context.Context) error {
			parse.Init("./local.yaml")
			// parse.Init("./config.yaml")
			hlog.Info("success init config")
			return nil
		})

		s.OnRun = append(s.OnRun, func(ctx context.Context) error {
			mw.InitJwt()
			hlog.Info("success init jwt")
			return nil
		})

		s.OnRun = append(s.OnRun, func(ctx context.Context) error {
			kitex_client.Init()
			hlog.Info("success init kitex client")
			return nil
		})

		s.OnRun = append(s.OnRun, func(ctx context.Context) error {
			mq.Init()
			mq.PublishSimple("favorite")
			mq.PublishSimple("follow")
			hlog.Info("success init mq")
			return nil
		})
	})
}

func ShutdownHook(s *server.Hertz) {
	shutdownOnce.Do(func() {
		s.OnShutdown = append(s.OnShutdown, func(ctx context.Context) {
			mq.Close()
			hlog.Info("success close mq")
		})
	})
}
