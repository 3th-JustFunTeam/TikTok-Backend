package dao

import "github.com/3th-JustFunTeam/Tiktok-Backend/model/po"

func SaveLike(love po.Love) error {

	tx := DB.Save(&love)
	return tx.Error
}

func DeleteLike(userId, videoId int) (error, int64) {

	var love = po.Love{}
	res := DB.Where("user_id = ? and video_id = ?", userId, videoId).Delete(&love)
	return res.Error, res.RowsAffected
}

func QueryIsLike(userId, videoId uint) int64 {
	var count int64
	DB.Model(&po.Love{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count)
	return count
}

func VideoLikeCount(videoId uint) int64 {
	var count int64
	DB.Model(&po.Love{}).Where("video_id = ?", videoId).Count(&count)
	return count
}

func QueryLikeByUserId(userId string) ([]po.Love, error) {
	var love []po.Love
	res := DB.Where("user_id = ?", userId).Find(&love)
	return love, res.Error
}
