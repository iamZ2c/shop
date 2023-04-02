package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/forms"
	"mxshop_api/good_web/global"
	"mxshop_api/good_web/proto"
	"net/http"
	"strconv"
)

func Str2Int(s string) int32 {
	c, err := strconv.Atoi(s)
	if err != nil {
		panic(c)
	}
	return int32(c)
}

func List(ctx *gin.Context) {
	req := proto.GoodsFilterReq{}
	pmin := ctx.DefaultQuery("pmin", "0")
	pmax := ctx.DefaultQuery("pmax", "0")
	isHot := ctx.DefaultQuery("is_hot", "0")
	isNew := ctx.DefaultQuery("is_new", "0")
	isTab := ctx.DefaultQuery("is_tab", "0")
	page := ctx.DefaultQuery("page", "0")
	pageSize := ctx.DefaultQuery("page_size", "0")
	keyword := ctx.DefaultQuery("keyword", "")
	brandId := ctx.DefaultQuery("brand_id", "0")
	categoryId := ctx.DefaultQuery("category_id", "0")
	req.PriceMin = Str2Int(pmin)
	req.PriceMin = Str2Int(pmax)
	req.Pages = Str2Int(page)
	req.PagePerNums = Str2Int(pageSize)
	req.KeyWords = keyword
	req.TopCategory = Str2Int(categoryId)
	req.Brand = Str2Int(brandId)
	if isHot == "1" {
		req.IsHot = true
	}
	if isNew == "1" {
		req.IsNew = true
	}
	if isTab == "1" {
		req.IsTab = true
	}
	resp, err := global.GoodsSrvClient.GoodsList(context.Background(), &req)
	if err != nil {
		global.Check(ctx, err, "[goods-srv-list]")
	}
	goodsList := make([]map[string]interface{}, 2)
	for _, good := range resp.Data {
		goods := make(map[string]interface{})
		goods["id"] = good.Id
		goods["goods_sn"] = good.GoodsSn
		goods["goods_brief"] = good.GoodsBrief
		goods["goods_name"] = good.Name
		goods["is_hot"] = good.IsHot
		goods["is_new"] = good.IsNew
		goods["shop_price"] = good.ShopPrice
		goods["market_price"] = good.MarketPrice
		goods["on_sale"] = good.OnSale
		goods["fav_num"] = good.FavNum
		goods["sold_num"] = good.SoldNum
		goods["click_num"] = good.ClickNum
		goods["images"] = good.Images
		goods["desc_images"] = good.DescImages
		goods["goods_front_image"] = good.GoodsFrontImage
		goodsList = append(goodsList, goods)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  goodsList,
		"total": resp.Total,
	})
	return

}

func New(ctx *gin.Context) {
	f := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&f); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "参数验证失败",
				"info":    err.Error(),
			},
		)
		return
	}
	srvResp, err := global.GoodsSrvClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:    f.Name,
		GoodsSn: f.GoodsSn,
		//Stocks:          0,
		MarketPrice:     f.MarketPrice,
		ShopPrice:       f.ShopPrice,
		GoodsBrief:      f.GoodsBrief,
		ShipFree:        f.ShipFree,
		Images:          f.Images,
		DescImages:      f.DescImages,
		GoodsFrontImage: f.FrontImage,
		CategoryId:      f.CategoryId,
		BrandId:         f.Brand,
	})
	global.Check(ctx, err, "asd")
	ctx.JSON(
		http.StatusOK,
		srvResp,
	)
}

func Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	SrvResp, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoReq{Id: int32(id)})

	goods := make(map[string]interface{})
	// #TODO 库存服务
	goods["id"] = SrvResp.Id
	goods["goods_sn"] = SrvResp.GoodsSn
	goods["goods_brief"] = SrvResp.GoodsBrief
	goods["goods_name"] = SrvResp.Name
	goods["is_hot"] = SrvResp.IsHot
	goods["is_new"] = SrvResp.IsNew
	goods["shop_price"] = SrvResp.ShopPrice
	goods["market_price"] = SrvResp.MarketPrice
	goods["on_sale"] = SrvResp.OnSale
	goods["fav_num"] = SrvResp.FavNum
	goods["sold_num"] = SrvResp.SoldNum
	goods["click_num"] = SrvResp.ClickNum
	goods["images"] = SrvResp.Images
	goods["desc_images"] = SrvResp.DescImages
	goods["goods_front_image"] = SrvResp.GoodsFrontImage

	ctx.JSON(
		http.StatusOK,
		goods,
	)
	return
}

func DeleteGoods(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}

	_, err = global.GoodsSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: int32(id)})
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
}

func Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	s := forms.GoodsForm{}
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(id),
		Name:            s.Name,
		GoodsSn:         s.GoodsSn,
		Stocks:          s.Stocks,
		CategoryId:      s.CategoryId,
		MarketPrice:     s.MarketPrice,
		ShopPrice:       s.ShopPrice,
		GoodsBrief:      s.GoodsBrief,
		Images:          s.Images,
		DescImages:      s.DescImages,
		ShipFree:        s.ShipFree,
		GoodsFrontImage: s.FrontImage,
		BrandId:         s.Brand,
	})
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
}

func UpdateGoodsStatus(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	s := forms.GoodsStatus{}
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(id),
		IsNew:  false,
		IsHot:  false,
		OnSale: false,
	})
	if err != nil {
		global.Check(ctx, err, "dasd")
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
}
