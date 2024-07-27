/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 11:15:07
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-19 15:56:50
 * @Description: 
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\hertz-server\handler\comment_api\comment_service.go
 */
// Code generated by hertz generator.

package comment_api

import (
	"context"

	comment_api "douyin/hertz-server/model/comment_api"
	"douyin/hertz-server/model/user_api"
	"douyin/hertz-server/pkg/kitex_client"
	"douyin/hertz-server/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentAction .
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment_api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(comment_api.CommentActionResp)

	// 获取用户的id
	rowUserID, ok := c.Get("user_id")
	if !ok {
		resp.StatusCode = 1
		resp.StatusMsg = "could not find user_id"
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := rowUserID.(int64)

	switch req.ActionType {
	case 1: // 发布评论
		if req.CommentText == nil || *req.CommentText == "" {
			resp.StatusCode = 1
			resp.StatusMsg = "could not send empty comment"
			c.JSON(consts.StatusOK, resp)
			return
		}

		// 数据库添加评论记录
		commentActionRpc, err := kitex_client.CommentActionRpc(ctx, kitex_client.PubComment, userID, req.VideoID, req.CommentText)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		// 查看评论用户的信息
		userListRpc, err := kitex_client.UserListRpc(ctx, []int64{userID})
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		// 视频评论数自增
		err = kitex_client.VideoInfoActionRpc(ctx, kitex_client.PubComment, req.VideoID)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		resp.Comment = &comment_api.Comment{
			ID: commentActionRpc.Comment.Id,
			User: &user_api.User{
				ID:              userListRpc.Users[0].Id,
				Name:            userListRpc.Users[0].Name,
				FollowCount:     userListRpc.Users[0].FollowCount,
				FollowerCount:   userListRpc.Users[0].FollowerCount,
				IsFollow:        false, // 不能关注自己
				Avatar:          userListRpc.Users[0].Avatar,
				BackgroundImage: userListRpc.Users[0].Background,
				Signature:       userListRpc.Users[0].Signature,
				TotalFavorited:  userListRpc.Users[0].TotalFavorited,
				WorkCount:       userListRpc.Users[0].TotalFavorited,
				FavoriteCount:   userListRpc.Users[0].FavoriteCount,
			},
			Content:    commentActionRpc.Comment.Content,
			CreateDate: commentActionRpc.Comment.CreateDate,
		}
		resp.StatusCode = 0
		resp.StatusMsg = "successfully publish comment"
	case 2:
		if req.CommentID == nil || *req.CommentID == 0 {
			resp.StatusCode = 1
			resp.StatusMsg = "empty comment id"
			c.JSON(consts.StatusOK, resp)
			return
		}

		// 在数据库删除评论记录
		if _, err = kitex_client.CommentActionRpc(ctx, kitex_client.DelComment, userID, req.VideoID, nil); err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "empty comment id"
			c.JSON(consts.StatusOK, resp)
			return
		}

		// 视频评论数自减
		err = kitex_client.VideoInfoActionRpc(ctx, kitex_client.DelComment, req.VideoID)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	default:
		resp.StatusCode = 1
		resp.StatusMsg = "Invalid action type"
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment_api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(comment_api.CommentListResp)

	// 获取视频的评论列表
	commentListRpc, err := kitex_client.CommentListRpc(ctx, req.VideoID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	align := len(commentListRpc.Comment)

	// 获取评论列表对应的用户信息
	userIDList := make([]int64, 0, len(commentListRpc.Comment))
	for _, v := range commentListRpc.Comment {
		userIDList = append(userIDList, v.UserId)
	}
	userListRpc, err := kitex_client.UserListRpc(ctx, userIDList)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	if align != len(userListRpc.Users) {
		resp.StatusCode = 1
		resp.StatusMsg = "invalid align user list"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 查看评论列表的用户是否被关注
	var isFollowList []bool
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		isFollowList = make([]bool, 0, align)
		for i := 0; i < align; i++ {
			isFollowList = append(isFollowList, false)
		}
	} else {
		userID, err := mw.TokenGetUserId(req.Token)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFollowRpc, err := kitex_client.IsFollowRpc(ctx, userID, userIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		if align != len(isFollowRpc.IsFollow) {
			resp.StatusCode = 1
			resp.StatusMsg = "invalid align user list"
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFollowList = isFollowRpc.IsFollow
	}

	for i := 0; i < align; i++ {
		resp.CommentList = append(resp.CommentList, &comment_api.Comment{
			ID: commentListRpc.Comment[i].Id,
			User: &user_api.User{
				ID:              userListRpc.Users[i].Id,
				Name:            userListRpc.Users[i].Name,
				FollowCount:     userListRpc.Users[i].FollowCount,
				FollowerCount:   userListRpc.Users[i].FollowerCount,
				IsFollow:        isFollowList[i],
				Avatar:          userListRpc.Users[i].Avatar,
				BackgroundImage: userListRpc.Users[i].Background,
				Signature:       userListRpc.Users[i].Signature,
				TotalFavorited:  userListRpc.Users[i].TotalFavorited,
				WorkCount:       userListRpc.Users[i].WorkCount,
				FavoriteCount:   userListRpc.Users[i].FavoriteCount,
			},
			Content:    commentListRpc.Comment[i].Content,
			CreateDate: commentListRpc.Comment[i].CreateDate,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}