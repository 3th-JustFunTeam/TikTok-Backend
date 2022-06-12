package service

import (
	"errors"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
)

func Like(token string, video_id int) error {
	_, err := utils.ParseToken(token)
	if err != nil {
		return errors.New("token 解析错误")
	}
	err = dao.LikeUp(uint(video_id))
	if err != nil {
		return errors.New("sql 运行错误")
	}
	return nil
}

func DisLike(video_id int) error {
	err := dao.LikeDown(uint(video_id))
	if err != nil {
		return errors.New("sql 运行错误")
	}
	return nil
}
