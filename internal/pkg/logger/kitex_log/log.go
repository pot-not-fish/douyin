/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-18 19:31:21
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-19 00:06:49
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\logger\kitex_log\log.go
 */
package kitex_log

import (
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)
}
