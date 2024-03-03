/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-12-12 16:30:21
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 00:14:52
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\pkg\dao\video_redis.go
 */
package dao

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

/**
 * @method Video
 * @description 将视频压入redis列表中，用于后续的视频feed流
 * @param
 * @return (error)
 */
func (video *Video) Publish() error {
	if video.ID <= 0 {
		return ErrInvalidUserId
	}

	redisDB.LPush("VideoId", video.ID)

	len, err := redisDB.LLen("VideoId").Result()
	if err != nil {
		return err
	}

	if len > 1000 {
		redisDB.RPop("VideoId")
	}

	if err != nil {
		return err
	}
	return nil
}

/**
 * @function
 * @description 获取redis视频列表中的某一段作为视频feed流
 * @param (offset int64) 决定从哪一段开始获取
 * @param (num int64) 获取的长度多少
 * @return ([]Video, int64, error) 返回视频列表，下一次从哪一段开始
 */
func VideoFeed(offset int64, limit int64) ([]Video, int64, error) {

	// 保证每次拿到的偏移量都能够整除
	if offset%limit != 0 {
		offset -= offset % limit
	}

	len, err := redisDB.LLen("VideoId").Result()
	if err != nil {
		return nil, 0, err
	}

	// 视频遍历完要循环从头开始
	if offset >= len {
		offset = 0
	}

	// 边界条件
	if offset+limit > len {
		limit = len - offset
	}

	str_videoid_list, err := redisDB.LRange("VideoId", offset, limit-1).Result()
	if err != nil {
		return nil, 0, err
	}

	var videoid_list []int64
	for _, v := range str_videoid_list {
		videoid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, 0, err
		}
		videoid_list = append(videoid_list, videoid)
	}

	videos, err := RetrieveVideos(videoid_list)
	if err != nil {
		return nil, 0, err
	}

	return videos, offset + limit, nil
}

/**
 * @method
 * @description 在缓存中创建视频信息
 * @param
 * @return
 */
func (v *Video) CreateVideoCache() error {
	var err error

	if v.ID <= 0 {
		return ErrInvalidVideoId
	}

	if v.PlayUrl == "" {
		return ErrEmptyPlayUrl
	}

	if v.CoverUrl == "" {
		return ErrEmptyCoverUrl
	}

	if v.UserID <= 0 {
		return ErrInvalidUserId
	}

	if v.Title == "" {
		return ErrEmptyTitle
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		id := fmt.Sprintf("%d", v.ID)

		err = pipe.HSet("video_play_url", id, v.PlayUrl).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet("video_cover_url", id, v.CoverUrl).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet("video_title", id, v.Title).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet("video_favorite_count", id, v.FavoriteCount).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet("video_comment_count", id, v.CommentCount).Err()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @method
 * @description 在缓存中查找视频信息
 * @param
 * @return
 */
func (v *Video) RetrieveVideoCache() error {
	var err error

	if v.ID <= 0 {
		return ErrInvalidVideoId
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		id := fmt.Sprintf("%d", v.ID)

		v.PlayUrl, err = pipe.HGet("video_play_url", id).Result()
		if err != nil {
			return err
		}

		v.CoverUrl, err = pipe.HGet("video_cover_url", id).Result()
		if err != nil {
			return err
		}

		v.Title, err = pipe.HGet("video_title", id).Result()
		if err != nil {
			return err
		}

		favorite_count, err := pipe.HGet("video_favorite_count", id).Result()
		if err != nil {
			return err
		}
		v.FavoriteCount, err = strconv.ParseInt(favorite_count, 10, 64)
		if err != nil {
			return err
		}

		comment_count, err := pipe.HGet("video_comment_count", id).Result()
		if err != nil {
			return err
		}
		v.CommentCount, err = strconv.ParseInt(comment_count, 10, 64)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户发布评论后，视频评论数自增1
 * @param
 * @return
 */
func IncCommentCache(video_id int64) error {
	var err error

	if err = redisDB.HIncrBy("video_comment_count", fmt.Sprintf("%d", video_id), 1).Err(); err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户删除评论后，视频评论数自减1
 * @param
 * @return
 */
func DecCommentCache(video_id int64) error {
	var err error

	if err = redisDB.HIncrBy("video_comment_count", fmt.Sprintf("%d", video_id), -1).Err(); err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户点赞后，视频点赞数自增1
 * @param
 * @return
 */
func IncFavoriteCache(video_id int64) error {
	var err error

	if err = redisDB.HIncrBy("video_favorite_count", fmt.Sprintf("%d", video_id), 1).Err(); err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户取消点赞后，视频点赞量自减1
 * @param
 * @return
 */
func DecFavoriteCache(video_id int64) error {
	var err error

	if err = redisDB.HIncrBy("video_favorite_count", fmt.Sprintf("%d", video_id), -1).Err(); err != nil {
		return err
	}

	return nil
}
