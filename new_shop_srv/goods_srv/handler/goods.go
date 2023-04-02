package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"new_shop_srv/goods_srv/global"
	"new_shop_srv/goods_srv/model"
	"new_shop_srv/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func Goods2Resp(goods model.Goods) proto.GoodsInfoResp {
	return proto.GoodsInfoResp{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		GoodsDesc:       "",
		ShipFree:        goods.ShipFree,
		Images:          goods.Images,
		DescImages:      goods.DescImages,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		AddTime:         0,
		Category:        nil,
		Brand:           nil,
	}

}

func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterReq) (*proto.GoodsListResp, error) {

	localDB := global.DB
	localDB = localDB.Model(&model.Goods{})
	if req.KeyWords != "" {
		localDB = localDB.Where("name LIKE '%?%'", req.KeyWords)
	}
	if req.IsNew {
		localDB = localDB.Where(&model.Goods{IsNew: true})
	}
	if req.IsHot {
		localDB = localDB.Where(&model.Goods{IsHot: true})
	}
	if req.PriceMax > 0 {
		localDB = localDB.Where("shop_price <= ?", req.PriceMax)
	}
	if req.PriceMin > 0 {
		localDB = localDB.Where("shop_price >= ?", req.PriceMin)
	}
	if req.TopCategory > 0 {
		var category model.Category
		global.DB.First(&category, req.TopCategory)
		if category.Level == 3 {
			localDB = localDB.Where("category_id=?", req.TopCategory)
		}
		if category.Level == 2 {
			localDB = localDB.Where(fmt.Sprintf("category_id IN (SELECT id FROM category WHERE parent_category_id = %v", req.TopCategory))
		}
		if category.Level == 1 {
			localDB = localDB.Where(fmt.Sprintf("category_id in (SELECT id FROM category WHERE parent_category_id in (SELECT id FROM category WHERE parent_category_id=%v))", req.TopCategory))
		}
	}
	var goods []model.Goods
	res := localDB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	data := make([]*proto.GoodsInfoResp, 2)
	for _, g := range goods {
		gResp := Goods2Resp(g)
		data = append(data, &gResp)
	}

	return &proto.GoodsListResp{
		Total: int32(res.RowsAffected),
		Data:  data,
	}, nil
}

func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResp, error) {
	// 购物车批量查询商品详情
	var goods []model.Goods
	res := global.DB.First(&goods, req.Id)
	data := make([]*proto.GoodsInfoResp, 2)
	for _, g := range goods {
		resp := Goods2Resp(g)
		data = append(data, &resp)
	}
	total := res.RowsAffected
	return &proto.GoodsListResp{
			Total: int32(total),
			Data:  data,
		},
		nil
}

func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResp, error) {
	var brand model.Brands
	res := global.DB.Find(&brand, req.BrandId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "没有相关分类")
	}
	var category model.Category
	res = global.DB.Find(&category, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "没有相关品牌")
	}
	goods := model.Goods{
		CategoryID:      req.CategoryId,
		BrandID:         req.BrandId,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		Name:            req.Name,
		GoodSn:          req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		Images:          nil,
		DescImages:      nil,
		GoodsFrontImage: "",
	}
	res = global.DB.Create(&goods)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, "创建失败")
	}
	return &proto.GoodsInfoResp{
			Id:   0,
			Name: ""},
		nil
}

func (s *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	var good model.Goods
	res := global.DB.Delete(&good)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, "创建失败")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	var brand model.Brands
	res := global.DB.Find(&brand, req.BrandId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "没有相关分类")
	}
	var category model.Category
	res = global.DB.Find(&category, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "没有相关品牌")
	}
	goods := model.Goods{
		CategoryID:      req.CategoryId,
		BrandID:         req.BrandId,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		Name:            req.Name,
		GoodSn:          req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		Images:          nil,
		DescImages:      nil,
		GoodsFrontImage: "",
	}
	res = global.DB.Save(&goods)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, "创建失败")
	}
	return &emptypb.Empty{},
		nil
}

//func (s *GoodsServer) GetGoodsDetail(context.Context, *proto.GoodInfoReq) (*proto.GoodsInfoResp, error) {
//}
