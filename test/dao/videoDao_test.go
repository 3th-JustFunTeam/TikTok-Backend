package dao

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"testing"
)

func Test_QueryAllVideo(t *testing.T) {
	dao.InitDB()
	video, err := dao.QueryAllVideo()

	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(video)
}
