package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID       int32     `gorm:"primarykey"`
	CreateAt time.Time `gorm:"column:add_time"`
	UpdateAt time.Time `gorm:"column:update_time"`
	DeleteAt gorm.DeletedAt
	IsDelete bool
}

/*
1.密文 2.密文不可反解
1.非对称加密
2.wdmd5
*/

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female女 male男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户，2表示管理员'"`
}
