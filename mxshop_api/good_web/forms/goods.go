package forms

type GoodsForm struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=100"`
	GoodsSn     string   `form:"goods_sn" json:"goods_sn" binding:"required,min=2,max=100"`
	Stocks      int32    `form:"stocks" json:"stocks" binding:"required,min=1"`
	CategoryId  int32    `form:"category_id" json:"category_id" binding:"required"`
	MarketPrice float32  `form:"market_price" json:"market_price" binding:"required"`
	ShopPrice   float32  `form:"shop_price" json:"shop_price" binding:"required"`
	GoodsBrief  string   `form:"goods_brief" json:"goods_brief" binding:"required"`
	Images      []string `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string `form:"desc_images" json:"desc_images" binding:"required,min=1"`
	ShipFree    bool     `form:"ship_free" json:"ship_free" binding:"required"`
	FrontImage  string   `form:"front_image" json:"front_image" binding:"required"`
	Brand       int32    `form:"brand" json:"brand" binding:"required"`
}

type GoodsStatus struct {
	IsNew  bool `form:"is_new" json:"is_new" binding:"required"`
	IsHot  bool `form:"is_hot" json:"is_hot" binding:"required"`
	OnSale bool `form:"on_sale" json:"on_sale" binding:"required"`
}
