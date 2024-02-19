/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-25 23:22:47
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2023-12-25 23:37:11
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\hertz-server\router\favorite_api\middleware.go
 */
// Code generated by hertz generator.

package favorite_api

import (
	"douyin/internal/pkg/logger/hertz_log"
	"douyin/internal/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{hertz_log.AccessLog()}
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteactionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc(), mw.SaveUserId}
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}
