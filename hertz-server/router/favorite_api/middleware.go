/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 11:15:02
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-02 12:24:42
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\hertz-server\router\favorite_api\middleware.go
 */
// Code generated by hertz generator.

package favorite_api

import (
	"douyin/hertz-server/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteactionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc(), mw.SaveUserId}
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}
