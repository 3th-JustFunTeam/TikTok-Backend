package po

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;type:bigint unsigned auto_increment;primaryKey;not null"`
	AuthName  string    `gorm:"column:auth_name;type:nvarchar(32);unique;not null"`
	NickName  string    `gorm:"column:nick_name;type:nvarchar(32);not null"`
	Password  string    `gorm:"column:nick_name;type:varchar(32);not null"`
	CreatedAt time.Time `gorm:"column:ts;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"column:update_ts;type:timestamp;not null"`
	IsDeleted bool      `gorm:"column:is_deleted;type:bool;not null"`
}
