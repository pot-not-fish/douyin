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
	"douyin/internal/pkg/dal/user_dal"
	"douyin/internal/pkg/dal/video_dal"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
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

	// 获取视频feed流信息
	videos, next_time, err := video_dal.VideoSubscribe(req.LatestTime, 10)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Invalid timestamp"
		c.JSON(consts.StatusOK, resp)
		return
	}

	if len(videos) == 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "nobody upload video up to now"
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取视频用户信息
	var video_user_id_slice []int64
	for _, v := range videos {
		video_user_id_slice = append(video_user_id_slice, v.UserId)
	}
	respRpc, err := kitex_client.UserinfoRpc(ctx, video_user_id_slice)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	var isfavorite_list = make([]bool, 0, 10)
	var isfollow_list = make([]bool, 0, 10)
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < 10; i++ {
			isfavorite_list = append(isfavorite_list, false)
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

		var userid_list = make([]int64, 0, 10)
		var videoid_list = make([]int64, 0, 10)
		var to_user_id_list = make([]int64, 0, 10)
		for _, v := range videos {
			videoid_list = append(videoid_list, int64(v.ID))
			userid_list = append(userid_list, user_id)
			to_user_id_list = append(to_user_id_list, v.UserId)
		}

		isfavorite_list, err = kitex_client.IsFavoriteRpc(ctx, userid_list, videoid_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		isfollow_list, err = kitex_client.IsFollowRpc(ctx, userid_list, to_user_id_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	}

	for k, v := range videos {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: int64(v.ID),
			Author: &user_api.User{
				ID:              respRpc.User[k].UserId,
				Name:            respRpc.User[k].Name,
				FollowCount:     respRpc.User[k].FollowCount,
				FollowerCount:   respRpc.User[k].FollowerCount,
				IsFollow:        isfollow_list[k],
				Avatar:          respRpc.User[k].Avatar,
				BackgroundImage: respRpc.User[k].Background,
				Signature:       respRpc.User[k].Signature,
				TotalFavorited:  respRpc.User[k].TotalFavorited,
				WorkCount:       respRpc.User[k].WorkCount,
				FavoriteCount:   respRpc.User[k].FavoriteCount,
			},
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + v.PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isfavorite_list[k],
			Title:         v.Title,
		})
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	resp.NextTime = next_time

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

	videos, err := video_dal.RetrieveUserVideos(req.UserID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 将id发送给kitex userinfo的rpc服务，获取用户信息
	userinfo_list, err := kitex_client.UserinfoRpc(ctx, []int64{req.UserID})
	if err != nil {
		resp.StatusMsg = err.Error()
		resp.StatusCode = 1
		c.JSON(consts.StatusOK, resp)
		return
	}

	var isfavorite_list = make([]bool, 0, len(videos))
	var isfollow_list = make([]bool, 0, len(videos))
	if req.Token == nil || *req.Token == "" { // nil对应不存在token字段，""对应token值为空
		for i := 0; i < len(videos); i++ {
			isfavorite_list = append(isfavorite_list, false)
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

		var user_id_list = make([]int64, 0, len(videos))
		var to_user_id_list = make([]int64, 0, len(videos))
		var videoid_list = make([]int64, 0, len(videos))
		for _, v := range videos {
			videoid_list = append(videoid_list, int64(v.ID))
			user_id_list = append(user_id_list, user_id)
			to_user_id_list = append(to_user_id_list, v.UserId)
		}

		isfavorite_list, err = kitex_client.IsFavoriteRpc(ctx, user_id_list, videoid_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}

		isfollow_list, err = kitex_client.IsFollowRpc(ctx, user_id_list, to_user_id_list)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			c.JSON(consts.StatusOK, resp)
			return
		}
	}

	for k, v := range videos {
		resp.VideoList = append(resp.VideoList, &video_api.Video{
			ID: int64(v.ID),
			Author: &user_api.User{
				ID:              userinfo_list.User[0].UserId,
				Name:            userinfo_list.User[0].Name,
				FollowCount:     userinfo_list.User[0].FollowCount,
				FollowerCount:   userinfo_list.User[0].FollowerCount,
				IsFollow:        isfollow_list[k],
				Avatar:          userinfo_list.User[0].Avatar,
				BackgroundImage: userinfo_list.User[0].Background,
				Signature:       userinfo_list.User[0].Signature,
				TotalFavorited:  userinfo_list.User[0].TotalFavorited,
				WorkCount:       userinfo_list.User[0].WorkCount,
				FavoriteCount:   userinfo_list.User[0].FavoriteCount,
			},
			PlayURL:       "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + v.PlayUrl,
			CoverURL:      "https://840231514-1320167793.cos.ap-nanjing.myqcloud.com" + v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isfavorite_list[k],
			Title:         v.Title,
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
	row_user_id, ok := c.Get("user_id")
	if !ok {
		resp.StatusCode = 1
		resp.StatusMsg = "could not find user_id"
		c.JSON(consts.StatusOK, resp)
		return
	}
	user_id := row_user_id.(int64)

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
	timestamp := time.Now().Unix()
	var video = video_dal.Video{
		CoverUrl: fmt.Sprintf("/src/picture/%d-%s-%d.jpg", user_id, req.Title, timestamp),
		PlayUrl:  fmt.Sprintf("/src/video/%d-%s-%d.mp4", user_id, req.Title, timestamp),
		UserId:   user_id,
		Title:    req.Title,
	}

	// 上传视频文件和视频封面文件
	if err = Uploadfile(ctx, dataFile, video.PlayUrl, video.CoverUrl); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	if err = video.CreateVideo(); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	var user = &user_dal.User{Model: gorm.Model{ID: uint(user_id)}}
	if err := user.IncWorkCount(); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 将视频推送到feed流中
	if err = video.Publish(); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "OK"
	c.JSON(consts.StatusOK, resp)
}

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
