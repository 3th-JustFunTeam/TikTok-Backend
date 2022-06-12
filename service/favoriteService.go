package service

import (
	"errors"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
)

func Like(user_id, video_id int) error {
	err := dao.LikeUp(uint(user_id), uint(video_id))
	if err != nil {
		return errors.New("sql 运行错误")
	}
	return nil
}

func DisLike(user_id, video_id int) error {
	err := dao.LikeDown(uint(user_id), uint(video_id))
	if err != nil {
		return errors.New("sql 运行错误")
	}
	return nil
}
