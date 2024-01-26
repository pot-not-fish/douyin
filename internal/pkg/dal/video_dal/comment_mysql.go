package video_dal

import "gorm.io/gorm"

/**
 * @method
 * @description 用户发布评论，创建相应数据库字段
 * @param
 * @return
 */
func (c *Comment) CreateComment() error {
	if VideoDb == nil {
		return ErrNullVideoDb
	}

	if c.Content == "" {
		return ErrEmptyComment
	}

	if c.UserId <= 0 {
		return ErrInvalidUserId
	}

	if c.VideoRefer == 0 {
		return ErrEmptyVideoId
	}

	create_comment(c)

	return nil
}

func create_comment(c *Comment) error {
	return VideoDb.Transaction(func(tx *gorm.DB) error {
		var video Video
		if err := tx.Where("id = ?", c.VideoRefer).First(&video).Error; err != nil {
			return err
		}

		if err := tx.Model(&video).Updates(map[string]interface{}{
			"comment_count": gorm.Expr("comment_count + ?", 1),
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&video).Association("Comments").Append(c); err != nil {
			return err
		}

		return nil
	})
}

/**
 * @method
 * @description 用户删除评论，删除comments表相关字段（软删除）
 * @param
 * @return
 */
func (c *Comment) DeleteComment() error {
	if VideoDb == nil {
		return ErrNullVideoDb
	}

	if c.ID == 0 {
		return ErrEmptyCommentId
	}

	if c.VideoRefer == 0 {
		return ErrEmptyVideoId
	}

	delete_comment(c)

	return nil
}

func delete_comment(c *Comment) error {
	return VideoDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("video_refer = ? AND id = ?", c.VideoRefer, c.ID).First(c).Error; err != nil {
			return err
		}

		if err := tx.Model(&Video{}).Where("video_refer = ?", c.VideoRefer).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Delete(c).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
 * @function
 * @description 查找视频的相关评论，按照时间顺序返回
 * @param
 * @return
 */
func RetrieveComment(video_id int64) ([]Comment, error) {
	if VideoDb == nil {
		return nil, ErrNullVideoDb
	}

	if video_id == 0 {
		return nil, ErrEmptyVideoId
	}

	var comments = make([]Comment, 0, 20)

	if err := VideoDb.Order("created_at desc").Where("video_refer = ?", video_id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

/**
 * @method
 * @description 查找视频相关的评论
 * @param
 * @return
 */
func (v *Video) RetrieveComment() error {
	if VideoDb == nil {
		return ErrNullVideoDb
	}

	if v.ID == 0 {
		return ErrEmptyVideoId
	}

	if err := VideoDb.Preload("Comments").Where("id = ?", v.ID).Find(v).Error; err != nil {
		return err
	}

	return nil
}
