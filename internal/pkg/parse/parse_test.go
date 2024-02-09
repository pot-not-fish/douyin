/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-09 17:34:50
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-09 19:48:49
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\parse\parse_test.go
 */
package parse

import "testing"

func TestInit(t *testing.T) {
	Init()
	t.Log(ConfigStructure)
}
