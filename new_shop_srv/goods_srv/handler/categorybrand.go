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

func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterReq) (*proto.CategoryBrandListResp, error) {
	var categoryBrands []*model.GoodsCategoryBrand
	res := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Preload("Category").Preload("Brands").Find(&categoryBrands)
	var count int64
	res.Count(&count)

	cbRespList := make([]*proto.CategoryBrandResp, 5)
	fmt.Println(cbRespList)

	for _, categorybrand := range categoryBrands {

		cbRespList = append(cbRespList, &proto.CategoryBrandResp{
			Id: categorybrand.ID,
			Brand: &proto.BrandInfoResp{
				Id:   categorybrand.BrandsID,
				Name: categorybrand.Brands.Name,
				Logo: categorybrand.Brands.Logo,
			},
			Category: &proto.CategoryInfoResp{
				Id:             categorybrand.CategoryID,
				Name:           categorybrand.Category.Name,
				ParentCategory: categorybrand.Category.ParentCategoryId,
				Level:          categorybrand.Category.Level,
				IsTab:          categorybrand.Category.IsTab,
			},
		})
	}

	resp := proto.CategoryBrandListResp{
		Total: int32(count),
		Data:  cbRespList,
	}
	return &resp, nil

}

func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoReq) (*proto.BrandListResp, error) {
	var categoryBrandList []model.GoodsCategoryBrand
	res := global.DB.Where("category_id=?", req.Id).Preload("Brands").Find(&categoryBrandList)
	var c int64
	res.Count(&c)
	biResp := make([]*proto.BrandInfoResp, 2)
	for _, categoryBrand := range categoryBrandList {
		biResp = append(biResp, &proto.BrandInfoResp{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}
	return &proto.BrandListResp{
		Total: int32(c),
		Data:  biResp,
	}, nil
}

// 1-16 04:55
//
func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandReq) (*proto.CategoryBrandResp, error) {
	var Category model.Category
	res := global.DB.Find(&Category, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有相关分类")
	}
	var Brand model.Brands
	res = global.DB.Find(&Brand, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有相关品牌")
	}
	var CategoryBrand model.GoodsCategoryBrand
	CategoryBrand.CategoryID = req.CategoryId
	CategoryBrand.BrandsID = req.BrandId
	CategoryBrand.ID = req.Id
	global.DB.Create(&CategoryBrand)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, res.Error.Error())
	}
	global.DB.Preload("Brand").Preload("Category").Find(&CategoryBrand, req.Id)
	return &proto.CategoryBrandResp{
		Id: req.Id,
		Brand: &proto.BrandInfoResp{
			Id:   CategoryBrand.Brands.ID,
			Name: CategoryBrand.Brands.Name,
			Logo: CategoryBrand.Brands.Logo,
		},
		Category: &proto.CategoryInfoResp{
			Id:             CategoryBrand.Category.ID,
			Name:           CategoryBrand.Category.Name,
			ParentCategory: CategoryBrand.Category.ParentCategoryId,
			Level:          CategoryBrand.Category.Level,
			IsTab:          CategoryBrand.Category.IsTab,
		},
	}, nil
}

func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandReq) (*emptypb.Empty, error) {
	var Category model.Category
	res := global.DB.Find(&Category, req.CategoryId)
	if res.RowsAffected != 0 {
		return nil, status.Error(codes.InvalidArgument, "请先删除相关分类")
	}
	var Brand model.Brands
	res = global.DB.Find(&Brand, req.CategoryId)
	if res.RowsAffected != 0 {
		return nil, status.Error(codes.InvalidArgument, "请先删除相关品牌")
	}
	var CategoryBrand model.GoodsCategoryBrand
	global.DB.Delete(&CategoryBrand, req.Id)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, res.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandReq) (*emptypb.Empty, error) {
	var Category model.Category
	res := global.DB.Find(&Category, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有相关分类")
	}
	var Brand model.Brands
	res = global.DB.Find(&Brand, req.CategoryId)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "没有相关品牌")
	}
	var CategoryBrand model.GoodsCategoryBrand
	CategoryBrand.CategoryID = req.CategoryId
	CategoryBrand.ID = req.Id
	CategoryBrand.BrandsID = req.BrandId
	global.DB.Save(&CategoryBrand)
	if res.Error != nil {
		return nil, status.Error(codes.Unavailable, res.Error.Error())
	}
	return &emptypb.Empty{}, nil
}
