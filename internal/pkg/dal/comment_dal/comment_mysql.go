/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-26 17:28:46
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-14 22:24:50
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\comment_dal\comment_mysql.go
 */
package comment_dal

import (
	"fmt"
)

var (
	// 评论内容不为空
	ErrEmptyContent = fmt.Errorf("empty content")

	// 非法用户id
	ErrInvalidUserID = fmt.Errorf("invalid user id")

	// 非法视频id
	ErrInvalidVideoID = fmt.Errorf("invalid video id")

	// 非法评论id
	ErrInvalidCommentID = fmt.Errorf("invalid comment id")

	// 无效删除评论，评论不存在
	ErrInvalidDeleteComment = fmt.Errorf("invalid delete comment")

	// 数据库指针未初始化
	ErrNullDB = fmt.Errorf("nullptr database")
)

/**
 * @method
 * @description 在数据库中创建评论
 * @param
 * @return
 */
func (c *Comment) CreateComment() error {
	var err error
	if CommentDb == nil {
		return ErrNullDB
	}

	if c.Content == "" {
		return ErrEmptyContent
	}

	if c.UserID <= 0 {
		return ErrInvalidUserID
	}

	if c.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = CommentDb.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) DeleteComment() error {
	var err error
	if CommentDb == nil {
		return ErrNullDB
	}

	if c.ID <= 0 {
		return ErrInvalidCommentID
	}

	if err = CommentDb.Where("id = ? AND user_id = ? AND video_id = ?", c.ID, c.UserID, c.VideoID).First(c).Error; err != nil {
		return ErrInvalidDeleteComment
	}

	if err = CommentDb.Where("id = ?", c.ID).Unscoped().Delete(c).Error; err != nil {
		return err
	}

	return nil
}

func RetrieveComment(video_id int64) ([]Comment, error) {
	if CommentDb == nil {
		return nil, ErrNullDB
	}

	if video_id <= 0 {
		return nil, ErrInvalidVideoID
	}

	comments := make([]Comment, 0, 20)
	CommentDb.Order("created_at desc").Limit(20).Where("video_id = ?", video_id).Find(&comments)

	return comments, nil
}
