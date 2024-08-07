/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-25 17:23:30
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 00:04:06
 * @Description: 提供了对video数据库封装后的操作函数
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\pkg\dao\video_mysql.go
 */
package dao

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrEmptyTitle = fmt.Errorf("video title is empty")

	ErrEmptyCoverUrl = fmt.Errorf("video cover url is empty")

	ErrEmptyPlayUrl = fmt.Errorf("video play url is empty")

	ErrInvalidUserId = fmt.Errorf("user id is invalid")

	ErrInvalidVideoId = fmt.Errorf("video id is empty")

	ErrInvalidFavorite = fmt.Errorf("you can't favorite or cancel favorite repeatly")

	ErrNullVideoDb = fmt.Errorf("video's db pointer is null")

	ErrEmptyComment = fmt.Errorf("comment is empty")

	ErrEmptyCommentId = fmt.Errorf("comment id is empty")
)

/**
 * @method {Video}
 * @description 用于创建一个数据库video字段，需要Title,CoverUrl,PlayUrl,UserId
 * @param
 * @return (error)
 */
func (video *Video) CreateVideo() error {
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if len(video.Title) == 0 {
		return ErrEmptyTitle
	}

	if len(video.CoverUrl) == 0 {
		return ErrEmptyCoverUrl
	}

	if len(video.PlayUrl) == 0 {
		return ErrEmptyPlayUrl
	}

	// raw_id := []byte(video.Title + video.UserId + time.Now().String())
	// hax := md5.Sum(raw_id)
	// video.VideoId = fmt.Sprintf("%x", hax)

	if err := videoDb.Create(&video).Error; err != nil {
		return err
	}

	if err := video.Publish(); err != nil {
		return err
	}

	go video.CreateVideoCache()

	return nil
}

/**
 * @function
 * @description 通过视频的id查找视频的信息
 * @param
 * @return
 */
func RetrieveVideos(videoid_list []int64) ([]Video, error) {
	var videos = make([]Video, 0, 10)

	if videoDb == nil {
		return nil, ErrNullVideoDb
	}

	if err := videoDb.Order("created_at desc").Where("id IN ?", videoid_list).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

/**
 * @function
 * @description 用于返回某个用户的视频
 * @param (user_id string)
 * @return ([]Video, error) 返回一个视频列表
 */
func RetrieveUserVideos(user_id int64) ([]Video, error) {
	if videoDb == nil {
		return nil, ErrNullVideoDb
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserId
	}

	var videos []Video
	if err := videoDb.Order("created_at desc").Where("user_id = ?", user_id).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func UpdateVideoCache(video_id int64) error {
	var err error
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if video_id <= 0 {
		return ErrInvalidVideoId
	}

	video := Video{Model: gorm.Model{ID: uint(video_id)}}
	if err = videoDb.Where("id = ?", video.ID).First(video).Error; err != nil {
		return err
	}

	go video.CreateVideoCache()

	return nil
}

func IncCommentCount(video_id int64) error {
	var err error
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if video_id <= 0 {
		return ErrInvalidVideoId
	}

	if err = IncCommentCache(video_id); err == nil {
		go incCommentCount(video_id)
		return nil
	}

	if err = incCommentCount(video_id); err != nil {
		return err
	}

	go UpdateVideoCache(video_id)

	return nil
}

func incCommentCount(video_id int64) error {
	if err := videoDb.Model(&Video{}).Where("id = ?", video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func DecCommentCount(video_id int64) error {
	var err error
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if video_id <= 0 {
		return ErrInvalidVideoId
	}

	if err = DecCommentCache(video_id); err == nil {
		go decCommentCount(video_id)
		return nil
	}

	if err = decCommentCount(video_id); err != nil {
		return err
	}

	go UpdateVideoCache(video_id)

	return nil
}

func decCommentCount(video_id int64) error {
	if err := videoDb.Model(&Video{}).Where("id = ?", video_id).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func IncFavoriteCount(video_id int64) error {
	var err error
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if video_id <= 0 {
		return ErrInvalidVideoId
	}

	if err = IncFavoriteCache(video_id); err == nil {
		go incFavoriteCount(video_id)
		return nil
	}

	if err = incFavoriteCount(video_id); err != nil {
		return err
	}

	go UpdateVideoCache(video_id)

	return nil
}

func incFavoriteCount(video_id int64) error {
	if err := videoDb.Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func DecFavoriteCount(video_id int64) error {
	var err error
	if videoDb == nil {
		return ErrNullVideoDb
	}

	if video_id <= 0 {
		return ErrInvalidVideoId
	}

	if err = DecFavoriteCache(video_id); err == nil {
		go decFavoriteCount(video_id)
		return nil
	}

	if err = decFavoriteCount(video_id); err != nil {
		return err
	}

	go UpdateVideoCache(video_id)

	return nil
}

func decFavoriteCount(video_id int64) error {
	if err := videoDb.Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
