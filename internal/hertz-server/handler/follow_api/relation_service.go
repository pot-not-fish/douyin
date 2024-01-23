// Code generated by hertz generator.

package follow_api

import (
	"context"

	follow_api "douyin/internal/hertz-server/model/follow_api"
	"douyin/internal/hertz-server/model/user_api"
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationAction .
// @router /douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follow_api.RelationActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(follow_api.RelationActionResp)

	row_user_id, ok := c.Get("user_id")
	if !ok {
		resp.StatusCode = 1
		resp.StatusMsg = "missing user_id"
		c.JSON(consts.StatusOK, resp)
		return
	}
	user_id := row_user_id.(int64)

	switch req.ActionType {
	case 1:
		var relation = &user_dal.Relation{
			FollowId:   req.ToUserID,
			FollowerId: user_id,
		}

		if err = relation.CreateRelation(); err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	case 2:
		var relation = &user_dal.Relation{
			FollowId:   req.ToUserID,
			FollowerId: user_id,
		}

		if err = relation.DeleteRelation(); err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	default:
		resp.StatusCode = 1
		resp.StatusMsg = "Invalid action type"
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

// RelationFollow .
// @router /douyin/relatioin/follow/list [GET]
func RelationFollow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follow_api.FollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(follow_api.FollowListResp)

	relation_id_list, err := user_dal.RetrieveFollows(req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if len(relation_id_list) == 0 { // 特判用户没有任何关注的情况
		resp.StatusCode = 0
		resp.StatusMsg = "OK"
		c.JSON(consts.StatusOK, resp)
		return
	}

	var follow_id_list = make([]int64, 0, len(relation_id_list))
	for _, v := range relation_id_list {
		follow_id_list = append(follow_id_list, v.FollowId)
	}

	var isfollow_list = make([]bool, 0, len(relation_id_list))
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < len(relation_id_list); i++ {
			isfollow_list = append(isfollow_list, false)
		}
	} else {
		user_id, err := mw.TokenGetUserId(req.Token)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		var user_id_list = make([]int64, 0, len(relation_id_list))
		for i := 0; i < len(relation_id_list); i++ {
			user_id_list = append(user_id_list, user_id)
		}

		isfollow_list, err = kitex_client.IsFollowRpc(ctx, user_id_list, follow_id_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	}

	userinfo_list, err := kitex_client.UserinfoRpc(ctx, follow_id_list)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	for k, v := range userinfo_list.User {
		resp.UserList = append(resp.UserList, &user_api.User{
			ID:              v.UserId,
			Name:            v.Name,
			FollowCount:     v.FollowCount,
			FollowerCount:   v.FollowerCount,
			IsFollow:        isfollow_list[k],
			Avatar:          v.Avatar,
			BackgroundImage: v.Background,
			Signature:       v.Signature,
			TotalFavorited:  v.TotalFavorited,
			WorkCount:       v.WorkCount,
			FavoriteCount:   v.FavoriteCount,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

// RelationFollower .
// @router /douyin/relation/follower/list [GET]
func RelationFollower(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follow_api.FollowerListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(follow_api.FollowerListResp)

	relation_id_list, err := user_dal.RetrieveFollowers(req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if len(relation_id_list) == 0 { // 特判用户没有任何关注的情况下
		resp.StatusCode = 0
		resp.StatusMsg = "OK"
		c.JSON(consts.StatusOK, resp)
		return
	}

	var follower_id_list = make([]int64, 0, len(relation_id_list))
	for _, v := range relation_id_list {
		follower_id_list = append(follower_id_list, v.FollowerId)
	}

	var isfollow_list = make([]bool, 0, len(relation_id_list))
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < len(relation_id_list); i++ {
			isfollow_list = append(isfollow_list, false)
		}
	} else {
		user_id, err := mw.TokenGetUserId(req.Token)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		var user_id_list = make([]int64, 0, len(relation_id_list))
		for i := 0; i < len(relation_id_list); i++ {
			user_id_list = append(user_id_list, user_id)
		}

		isfollow_list, err = kitex_client.IsFollowRpc(ctx, user_id_list, follower_id_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	}

	userinfo_list, err := kitex_client.UserinfoRpc(ctx, follower_id_list)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	for k, v := range userinfo_list.User {
		resp.UserList = append(resp.UserList, &user_api.User{
			ID:              v.UserId,
			Name:            v.Name,
			FollowCount:     v.FollowCount,
			FollowerCount:   v.FollowerCount,
			IsFollow:        isfollow_list[k],
			Avatar:          v.Avatar,
			BackgroundImage: v.Background,
			Signature:       v.Signature,
			TotalFavorited:  v.TotalFavorited,
			WorkCount:       v.WorkCount,
			FavoriteCount:   v.FavoriteCount,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}
