package dao

import "github.com/3th-JustFunTeam/Tiktok-Backend/model/po"

func CreateFollowing(follow po.Follow) error {

	tx := DB.Create(&follow)
	return tx.Error

}

func DeleteFollowing(followingId string, userId int) error {

	var follow = po.Follow{}
	tx := DB.Where("user_id = ? and following_id = ?", userId, followingId).Delete(&follow)
	return tx.Error
}

func QueryFollowingListByUserId(userId int) ([]po.Follow, error) {
	var follow []po.Follow
	tx := DB.Where("user_id = ?", userId).Find(&follow)
	return follow, tx.Error
}

func QueryIsFollow(userId, followId int) int64 {
	var count int64
	DB.Model(&po.Follow{}).Where("user_id = ? and following_id = ?", userId, followId).Count(&count)
	return count
}

func QueryFollowingListByUserIdCount(userId int) (uint64, uint64) {
	var count int64
	var fensiCount int64
	DB.Model(&po.Follow{}).Where("user_id = ?", userId).Count(&count)
	DB.Model(&po.Follow{}).Where("following_id = ?", userId).Count(&fensiCount)

	return uint64(count), uint64(fensiCount)

}

func QueryFenshiListByUserId(userId int) ([]po.Follow, error) {
	var follow []po.Follow
	tx := DB.Where("following_id = ?", userId).Find(&follow)
	return follow, tx.Error
}
