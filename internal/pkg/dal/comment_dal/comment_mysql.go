package comment_dal

import (
	"fmt"

	"gorm.io/gorm"
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
)

/**
 * @method
 * @description 在数据库中创建评论
 * @param
 * @return
 */
func (c *Comment) CreateComment() error {
	var err error
	if c.Content == "" {
		return ErrEmptyContent
	}

	if c.UserID <= 0 {
		return ErrInvalidUserID
	}

	if c.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = create_comment(c); err != nil {
		return err
	}
	return nil
}

func create_comment(c *Comment) error {
	return CommentDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(c).Error; err != nil {
			return err
		}

		if err := tx.Model(&Video{}).Where("video_id = ?", c.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *Comment) DeleteComment() error {
	var err error
	if c.ID <= 0 {
		return ErrInvalidCommentID
	}

	if err = delete_comment(c); err != nil {
		return err
	}
	return nil
}

func delete_comment(c *Comment) error {
	return CommentDb.Transaction(func(tx *gorm.DB) error {
		// 只能删除自己的评论，指定某个视频的某个评论
		if err := tx.Where("id = ? AND user_id = ? AND video_id = ?", c.ID, c.UserID, c.VideoID).First(c).Error; err != nil {
			return ErrInvalidDeleteComment
		}

		if err := tx.Where("id = ?", c.ID).Unscoped().Delete(c).Error; err != nil {
			return err
		}

		if err := tx.Model(&Video{}).Where("video_id = ?", c.VideoID).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
