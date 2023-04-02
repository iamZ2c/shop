package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"new_shop_srv/goods_srv/global"
	"new_shop_srv/goods_srv/model"
	"new_shop_srv/goods_srv/proto"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//品牌和轮播图
func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterReq) (*proto.BrandListResp, error) {
	//var brands
	brands := make([]model.Brands, 2)
	fmt.Println(req)
	tx := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	// 统计记录条数
	var c int64
	global.DB.Model(&brands).Count(&c)
	fmt.Println(c)

	if tx.Error != nil {
		panic(tx.Error)
	}
	resp := &proto.BrandListResp{}
	var BrandList []*proto.BrandInfoResp

	for _, value := range brands {
		BrandList = append(BrandList,
			&proto.BrandInfoResp{
				Id:   value.ID,
				Name: value.Name,
				Logo: value.Logo,
			},
		)
	}
	resp.Total = int32(tx.RowsAffected)
	resp.Data = BrandList
	return resp, nil
}

// CreateBrand 添加商品，不能添加重复的品牌
func (s *GoodsServer) CreateBrand(ctx context.Context, br *proto.BrandReq) (*proto.BrandInfoResp, error) {
	var brand model.Brands

	res := global.DB.Where("name=?", br.Name).First(&brand)
	if res.RowsAffected != 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌已存在")
	}
	brand.ID = br.Id
	brand.Name = br.Name
	brand.Logo = br.Logo
	res = global.DB.Create(&brand)
	if res.Error != nil {
		return nil, status.Error(codes.InvalidArgument, res.Error.Error())
	}
	return &proto.BrandInfoResp{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo,
	}, nil
}

// DeleteBrand 删除商品，未删除的时候报错
func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandReq) (*emptypb.Empty, error) {
	var brand model.Brands
	res := global.DB.Delete(&brand, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandReq) (*emptypb.Empty, error) {
	var brand model.Brands
	res := global.DB.First(&brand)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌不存在")
	}

	brand.Name = req.Name
	brand.Logo = req.Logo
	res = global.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}
