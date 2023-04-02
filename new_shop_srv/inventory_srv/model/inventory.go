package model

type Inventory struct {
	BaseModel
	GoodsId int32 `gorm:"type:int"`
	// 库存
	Stocks int32 `gorm:"type:int"`
	// 分布式锁乐观锁
	Version int32 `gorm:"type:int"`
}
