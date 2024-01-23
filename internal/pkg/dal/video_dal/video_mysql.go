/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-25 17:23:30
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-21 19:42:38
 * @Description: 提供了对video数据库封装后的操作函数
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\video_dal\video_mysql.go
 */
package video_dal

import (
	"fmt"
)

var (
	ErrEmptyTitle = fmt.Errorf("video title is empty")

	ErrEmptyCoverUrl = fmt.Errorf("video cover url is empty")

	ErrEmptyPlayUrl = fmt.Errorf("video play url is empty")

	ErrInvalidUserId = fmt.Errorf("user id is invalid")

	ErrEmptyVideoId = fmt.Errorf("video id is empty")

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
	if VideoDb == nil {
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

	if err := VideoDb.Create(&video).Error; err != nil {
		return err
	}

	return nil
}

func RetrieveVideos(videoid_list []int64) ([]Video, error) {
	var videos = make([]Video, 0, 10)

	if VideoDb == nil {
		return nil, ErrNullVideoDb
	}

	if err := VideoDb.Order("created_at desc").Where("id IN ?", videoid_list).Find(&videos).Error; err != nil {
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
	if VideoDb == nil {
		return nil, ErrNullVideoDb
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserId
	}

	var videos []Video
	if err := VideoDb.Order("created_at desc").Where("user_id = ?", user_id).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func RetrieveUser(video_id int64) (int64, error) {
	var video Video

	if VideoDb == nil {
		return 0, ErrNullVideoDb
	}

	if video_id == 0 {
		return 0, ErrEmptyVideoId
	}

	if err := VideoDb.Where("id = ?", video_id).First(&video).Error; err != nil {
		return 0, err
	}

	return video.UserId, nil
}
