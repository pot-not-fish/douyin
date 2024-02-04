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

	videoList, err := kitex_client.VideoFeedRpc(ctx, req.LatestTime)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if len(videoList.Videos) == 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "nobody upload video up to now"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取视频用户信息
	var videoUserID []int64
	for _, v := range videoList.Videos {
		videoUserID = append(videoUserID, v.Id)
	}
	userList, err := kitex_client.UserListRpc(ctx, videoUserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	isFavoriteList := make([]bool, 0, len(userList.Users))
	isFollowList := make([]bool, 0, len(userList.Users))
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < len(userList.Users); i++ {
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

		videoIDList := make([]int64, 0, len(userList.Users))
		toUserIDList := make([]int64, 0, len(userList.Users))
		for _, v := range videoList.Videos {
			videoIDList = append(videoIDList, int64(v.Id))
			toUserIDList = append(toUserIDList, v.UserId)
		}

		isFavoriteRpc, err := kitex_client.IsFavoriteRpc(ctx, userID, videoIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFavoriteList = isFavoriteRpc.IsFavorite

		isFollowRpc, err := kitex_client.IsFollowRpc(ctx, userID, toUserIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFavoriteList = isFollowRpc.IsFollow
	}

	for i := 0; i < len(videoList.Videos); i++ {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: videoList.Videos[i].Id,
			Author: &user_api.User{
				ID:              userList.Users[i].Id,
				Name:            userList.Users[i].Name,
				FollowCount:     userList.Users[i].FollowCount,
				FollowerCount:   userList.Users[i].FollowerCount,
				IsFollow:        isFavoriteList[i],
				Avatar:          userList.Users[i].Avatar,
				BackgroundImage: userList.Users[i].Background,
				Signature:       userList.Users[i].Signature,
				TotalFavorited:  userList.Users[i].TotalFavorited,
				WorkCount:       userList.Users[i].WorkCount,
				FavoriteCount:   userList.Users[i].FavoriteCount,
			},
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoList.Videos[i].PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoList.Videos[i].CoverUrl,
			FavoriteCount: videoList.Videos[i].FavoriteCount,
			CommentCount:  videoList.Videos[i].CommentCount,
			IsFavorite:    isFavoriteList[i],
			Title:         videoList.Videos[i].Title,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	resp.NextTime = videoList.NextOffset
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

	videoList, err := kitex_client.VideoListRpc(ctx, req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 将id发送给kitex userinfo的rpc服务，获取用户信息
	userList, err := kitex_client.UserListRpc(ctx, []int64{req.UserID})
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = 1
		c.JSON(consts.StatusOK, resp)
		return
	}

	var isFavoriteList = make([]bool, 0, len(userList.Users))
	var isFollowList = make([]bool, 0, len(userList.Users))
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < len(userList.Users); i++ {
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

		var toUserIDList = make([]int64, 0, len(userList.Users))
		var videoIDList = make([]int64, 0, len(userList.Users))
		for i := 0; i < len(userList.Users); i++ {
			toUserIDList = append(toUserIDList, videoList.Videos[i].UserId)
			videoIDList = append(videoIDList, videoList.Videos[i].Id)
		}

		isFavoriteRpc, err := kitex_client.IsFavoriteRpc(ctx, userID, videoIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFavoriteList = isFavoriteRpc.IsFavorite

		isFollowRpc, err := kitex_client.IsFollowRpc(ctx, userID, toUserIDList)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
		isFollowList = isFollowRpc.IsFollow
	}

	for i := 0; i < len(userList.Users); i++ {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: videoList.Videos[i].Id,
			Author: &user_api.User{
				ID:              userList.Users[i].Id,
				Name:            userList.Users[i].Name,
				FollowCount:     userList.Users[i].FollowCount,
				FollowerCount:   userList.Users[i].FollowerCount,
				IsFollow:        isFollowList[i],
				Avatar:          userList.Users[i].Avatar,
				BackgroundImage: userList.Users[i].Background,
				Signature:       userList.Users[i].Signature,
				TotalFavorited:  userList.Users[i].TotalFavorited,
				WorkCount:       userList.Users[i].WorkCount,
				FavoriteCount:   userList.Users[i].FavoriteCount,
			},
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoList.Videos[i].PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + videoList.Videos[i].CoverUrl,
			FavoriteCount: videoList.Videos[i].FavoriteCount,
			CommentCount:  videoList.Videos[i].CommentCount,
			IsFavorite:    isFavoriteList[i],
			Title:         videoList.Videos[i].Title,
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
	playURL := fmt.Sprintf("/src/picture/%d-%s-%d.jpg", userID, req.Title, timeStamp)
	coverURL := fmt.Sprintf("/src/video/%d-%s-%d.mp4", userID, req.Title, timeStamp)

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

	if err = kitex_client.UserInfoActionRpc(ctx, 3, userID, nil); err != nil {
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
	u, _ := url.Parse("https://840231514-1320167793.cos.ap-nanjing.myqcloud.com")

	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")

	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "",
			SecretKey: "",
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
