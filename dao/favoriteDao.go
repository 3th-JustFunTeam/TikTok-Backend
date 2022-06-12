package dao

import (
	"errors"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"gorm.io/gorm"
)

// 投稿后的视频点赞数存入数据库
func SaveLike(favorite po.Favorite) error {
	tx := DB.Create(&favorite)
	return tx.Error
}

// 判断是否已点赞
func IfLike(user_id, video_id uint) bool {
	var count int64
	DB.Model(&po.Favorite{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	return count > 0
}

// 点赞数+1
func LikeUp(user_id, video_id uint) error {
	tx := DB.Begin()

	// 判断是否点赞
	if IfLike(user_id, video_id) {
		tx.Rollback()
		return errors.New("点赞失败")
	}

	// 数据存入favorites表
	f := &po.Favorite{
		Video_Id:    video_id,
		User_Id:     user_id,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err := tx.Save(&f)
	if err != nil {
		tx.Rollback()
		return errors.New("sql favorites 储存错误")
	}

	// 数据存入Video表
	err = tx.Model(&po.Video{}).Where("video_id = ?", video_id).UpdateColumn("likecount", gorm.Expr("likecount + ?", 1))
	if err != nil {
		tx.Rollback()
		return errors.New("sql likeinfos 储存错误")
	}

	// 结束
	tx.Commit()
	return nil
}

// 点赞数-1
func LikeDown(user_id, video_id uint) error {
	tx := DB.Begin()

	// 判断是否点赞
	if !IfLike(user_id, video_id) {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}

	// favorites表中删除数据
	err := tx.Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&po.Favorite{})
	if err != nil {
		tx.Rollback()
		return errors.New("sql favorites 删除失败")
	}

	// Video表中减少数据
	err = tx.Model(&po.Video{}).Where("video_id = ?", video_id).UpdateColumn("likecount", gorm.Expr("likecount - ?", 1))
	if err != nil {
		tx.Rollback()
		return errors.New("sql likeinfos 减少失败")
	}

	// 结束
	tx.Commit()
	return nil
}
