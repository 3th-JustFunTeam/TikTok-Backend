package po

import "time"

type Favorite struct {
	Id          uint      `gorm:"column:id;autoIncrement"`
	Video_Id    uint      `gorm:"column:视频id"`
	User_Id     uint      `gorm:"column:用户id"`
	CreatedTime time.Time `gorm:"column:创建时间"`
	UpdatedTime time.Time `gorm:"column:修改时间"`
}
