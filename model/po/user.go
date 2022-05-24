package po

import "time"

type User struct {
	ID              uint64    `gorm:"column:id;type:bigint unsigned auto_increment;primaryKey;not null;comment:用户id"`
	AuthName        string    `gorm:"column:auth_name;type:nvarchar(32);unique;not null;comment:用户名"`
	NickName        string    `gorm:"column:nick_name;type:nvarchar(32);not null;comment:用户昵称"`
	Password        string    `gorm:"column:password;type:varchar(32);not null;comment:密码"`
	Avatar          string    `gorm:"column:avatar;type:varchar(100);not null;comment:用户头像地址"`
	Signature       string    `gorm:"column:signature;type:varchar(32);not null;comment:个性签名"`
	BackgroundImage string    `gorm:"column:background_image;type:varchar(100);not null;comment:背景图地址"`
	CreatedAt       time.Time `gorm:"column:ts;type:timestamp;not null;comment:创建时间"`
	UpdatedAt       time.Time `gorm:"column:update_ts;type:timestamp;not null;comment:修改时间"`
	IsDeleted       bool      `gorm:"column:is_deleted;type:bool;not null;comment:是否删除"`
}
