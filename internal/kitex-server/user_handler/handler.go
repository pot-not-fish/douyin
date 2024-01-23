/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-13 10:37:00
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-19 16:31:14
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\kitex-server\user_handler\handler.go
 */
package user_handler

import (
	"context"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_gen/user_rpc"

	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Userinfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) Userinfo(ctx context.Context, request *user_rpc.RetriveUserReq) (resp *user_rpc.RetriveUserResp, err error) {
	// TODO: Your code here...
	var resp_user []*user_rpc.User

	if len(request.UserId) == 0 {
		return &user_rpc.RetriveUserResp{
			StatusCode: 1,
			StatusMsg:  "empty user id slice",
		}, nil
	}

	for _, v := range request.UserId {
		var user = &user_dal.User{Model: gorm.Model{ID: uint(v)}}
		if err := user.RetrieveUser(); err != nil {
			return &user_rpc.RetriveUserResp{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}, nil
		}

		resp_user = append(resp_user, &user_rpc.User{
			UserId:         v,
			Name:           user.Name,
			FollowCount:    user.FollowCount,
			FollowerCount:  user.FollowerCount,
			Avatar:         user.Avatar,
			Background:     user.Background,
			Signature:      user.Signature,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		})
	}

	return &user_rpc.RetriveUserResp{
		StatusCode: 0,
		StatusMsg:  "OK",
		User:       resp_user,
	}, nil
}
