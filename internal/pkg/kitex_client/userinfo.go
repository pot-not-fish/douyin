/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-13 19:45:41
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-17 12:05:37
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\kitex_client\userinfo.go
 */
package kitex_client

import (
	"context"
	"douyin/internal/pkg/kitex_gen/user_rpc"
	"fmt"
)

func UserinfoRpc(ctx context.Context, userid []int64) (*user_rpc.RetriveUserResp, error) {
	respRpc, err := userinfoClient.Userinfo(ctx, &user_rpc.RetriveUserReq{UserId: userid})
	if err != nil {
		return nil, err
	}

	if respRpc.StatusCode != 0 {
		return nil, fmt.Errorf(respRpc.StatusMsg)
	}

	return respRpc, nil
}
