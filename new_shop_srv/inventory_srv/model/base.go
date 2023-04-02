package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GormList []string

// Scan 取值的时候因为只会返回error，所以这里需要引用传递，入的时候json加密会返回给value调用方
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), g)
}

// Value 自定义变量存入数据库方法
func (g GormList) Value() (value driver.Value, err error) {
	return json.Marshal(g)
}

type BaseModel struct {
	ID       int32     `gorm:"primarykey; type:int"`
	CreateAt time.Time `gorm:"column:add_time"`
	UpdateAt time.Time `gorm:"column:update_time"`
	DeleteAt gorm.DeletedAt
	IsDelete bool
}
