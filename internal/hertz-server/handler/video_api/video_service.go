// Code generated by hertz generator.

package video_api

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"douyin/internal/hertz-server/model/user_api"
	video_api "douyin/internal/hertz-server/model/video_api"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/mw"
	"douyin/internal/pkg/parse"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video_api.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video_api.FeedResp)

	// 时间戳不能是非法的请求
	if req.LatestTime < 0 {
		req.LatestTime = 0
	}

	// 获取视频feed流
	videoListRpc, err := kitex_client.VideoFeedRpc(ctx, req.LatestTime)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	align := len(videoListRpc.Videos)
	if align == 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "invalid align video list"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取视频用户信息
	var videoUserID []int64
	for _, v := range videoListRpc.Videos {
		videoUserID = append(videoUserID, v.UserId)
	}
	userListRpc, err := kitex_client.UserListRpc(ctx, videoUserID)
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

	var isFavoriteList []bool
	var isFollowList []bool
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		isFavoriteList = make([]bool, 0, align)
		isFollowList = make([]bool, 0, align)
		for i := 0; i < align; i++ {
			isFavoriteList = append(isFavoriteList, false)
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

		videoIDList := make([]int64, 0, align)
		toUserIDList := make([]int64, 0, align)
		for _, v := range videoListRpc.Videos {
			videoIDList = append(videoIDList, int64(v.Id))
			toUserIDList = append(toUserIDList, v.UserId)
		}

		// 查看视频是否点赞
		isFavoriteRpc, err := kitex_client.IsFavoriteRpc(ctx, userID, videoIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFavoriteList = isFavoriteRpc.IsFavorite

		// 查看用户是否关注
		isFollowRpc, err := kitex_client.IsFollowRpc(ctx, userID, toUserIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFollowList = isFollowRpc.IsFollow
	}

	for i := 0; i < align; i++ {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: videoListRpc.Videos[i].Id,
			Author: &user_api.User{
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
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoListRpc.Videos[i].PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoListRpc.Videos[i].CoverUrl,
			FavoriteCount: videoListRpc.Videos[i].FavoriteCount,
			CommentCount:  videoListRpc.Videos[i].CommentCount,
			IsFavorite:    isFavoriteList[i],
			Title:         videoListRpc.Videos[i].Title,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	resp.NextTime = videoListRpc.NextOffset
	c.JSON(consts.StatusOK, resp)
}

// List .
// @router /douyin/publish/list [GET]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video_api.ListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video_api.ListResp)

	// 获取用户发布的所有视频
	videoListRpc, err := kitex_client.VideoListRpc(ctx, req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	align := len(videoListRpc.Videos)
	if align <= 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "ok"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 将id发送给kitex userinfo的rpc服务，获取用户信息
	userListRpc, err := kitex_client.UserListRpc(ctx, []int64{req.UserID})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	var isFavoriteList []bool
	isFollow := false
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		isFavoriteList = make([]bool, 0, align)
		for i := 0; i < align; i++ {
			isFavoriteList = append(isFavoriteList, false)
		}
	} else {
		userID, err := mw.TokenGetUserId(req.Token)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		var videoIDList = make([]int64, 0, align)
		toUserID := videoListRpc.Videos[0].UserId
		for i := 0; i < align; i++ {
			videoIDList = append(videoIDList, videoListRpc.Videos[i].Id)
		}

		isFavoriteRpc, err := kitex_client.IsFavoriteRpc(ctx, userID, videoIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		if align != len(isFavoriteRpc.IsFavorite) {
			resp.StatusCode = 1
			resp.StatusMsg = "invalid align isfavorite list"
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFavoriteList = isFavoriteRpc.IsFavorite

		isFollowRpc, err := kitex_client.IsFollowRpc(ctx, userID, []int64{toUserID})
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFollow = isFollowRpc.IsFollow[0]
	}

	for i := 0; i < align; i++ {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: videoListRpc.Videos[i].Id,
			Author: &user_api.User{
				ID:              userListRpc.Users[0].Id,
				Name:            userListRpc.Users[0].Name,
				FollowCount:     userListRpc.Users[0].FollowCount,
				FollowerCount:   userListRpc.Users[0].FollowerCount,
				IsFollow:        isFollow,
				Avatar:          userListRpc.Users[0].Avatar,
				BackgroundImage: userListRpc.Users[0].Background,
				Signature:       userListRpc.Users[0].Signature,
				TotalFavorited:  userListRpc.Users[0].TotalFavorited,
				WorkCount:       userListRpc.Users[0].WorkCount,
				FavoriteCount:   userListRpc.Users[0].FavoriteCount,
			},
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoListRpc.Videos[i].PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoListRpc.Videos[i].CoverUrl,
			FavoriteCount: videoListRpc.Videos[i].FavoriteCount,
			CommentCount:  videoListRpc.Videos[i].CommentCount,
			IsFavorite:    isFavoriteList[i],
			Title:         videoListRpc.Videos[i].Title,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

type PublishResp struct {
	StatusCode int32  `thrift:"status_code,1" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `thrift:"status_msg,2" form:"status_msg" json:"status_msg" query:"status_msg"`
}

type VideoReq struct {
	Title string `form:"title" json:"title" query:"title" vd:"(len($) > 0 && len($) < 20); msg:'Illegal format'"`
}

// Publish .
// @router /douyin/publish [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	// 绑定title
	var req VideoReq // 如果用req := new(VideoReq) 后面&去掉

	resp := new(PublishResp)

	if err := c.BindAndValidate(&req); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "could not find title"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 在token中获取用户的user_id
	rowUserID, ok := c.Get("user_id")
	if !ok {
		resp.StatusCode = 1
		resp.StatusMsg = "could not find user_id"
		c.JSON(consts.StatusOK, resp)
		return
	}
	userID := rowUserID.(int64)

	// 获取视频文件
	dataFile, err := c.FormFile("data")
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 视频路径为/src/video/用户id-标题-时间戳
	// 图片路径为/src/picture/用户id-标题-时间戳
	timeStamp := time.Now().Unix()
	playURL := fmt.Sprintf("/src/picture/%d-%s-%d.mp4", userID, req.Title, timeStamp)
	coverURL := fmt.Sprintf("/src/video/%d-%s-%d.jpg", userID, req.Title, timeStamp)

	if err = Uploadfile(ctx, dataFile, playURL, coverURL); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if err = kitex_client.VideoActionRpc(ctx, userID, req.Title, coverURL, playURL); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if err = kitex_client.UserInfoActionRpc(ctx, kitex_client.IncUserWorkCount, userID, nil); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

/**
 * @function
 * @description 上传文件到腾讯cos对象存储
 * @param
 * @return
 */
func Uploadfile(ctx context.Context, dataFile *multipart.FileHeader, playurl string, coverurl string) error {
	u, _ := url.Parse(parse.ConfigStructure.Cos.BucketURL)

	su, _ := url.Parse(parse.ConfigStructure.Cos.ServiceURL)

	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  parse.ConfigStructure.Cos.SecretID,
			SecretKey: parse.ConfigStructure.Cos.SecretKey,
		},
	})

	fd, err := dataFile.Open()
	if err != nil {
		fmt.Println("Error opening video file")
		return err
	}

	_, err = client.Object.Put(ctx, playurl, fd, nil)
	if err != nil {
		fmt.Println("Error putting video file to cloud storage")
		return err
	}

	opt := &cos.GetSnapshotOptions{
		Time: 1, // 截取视频的第一秒
	}

	resp, err := client.CI.GetSnapshot(ctx, playurl, opt)
	if err != nil {
		fmt.Println("Error getting snapshot")
		return err
	}

	_, err = client.Object.Put(ctx, coverurl, resp.Body, nil)
	if err != nil {
		fmt.Println("Error putting snapshot to cloud storage")
		return err
	}

	return nil
}
