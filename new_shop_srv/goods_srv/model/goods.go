package model

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	// 为什么使用int32？不使用int，为了减少proto文件的转换
	Level            int32 `gorm:"type:int;not null;default:1;comment '目录等级'" json:"level"`
	ParentCategoryId int32 `json:"parent_category_id"`
	// 设置外键和ref,可以使用预加载的模式查询子品牌，不会生成字段
	SubCategory []Category `gorm:"foreignKey:ParentCategoryId;references:ID" json:"sub_category"`
	IsTab       bool       `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);not null;default:' '"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32
	Category   Category
	BrandsID   int32
	Brands     Brands
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;not null"`
	BrandID    int32 `gorm:"type:int;not null"`

	OnSale   bool `gorm:"default:false;not null;comment:'是否上架 '"`
	ShipFree bool `gorm:"default:false;not null;comment:'是否免邮'"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodSn          string   `gorm:"type:varchar(50);not null;comment:'商家用于查找物品编号'"`
	ClickNum        int32    `gorm:"type:int;not null;default:0"`
	SoldNum         int32    `gorm:"type:int;not null;default:0"`
	FavNum          int32    `gorm:"type:int;not null;default:0"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`
	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:'商品标题图片'"`
}
