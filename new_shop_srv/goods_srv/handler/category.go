package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"new_shop_srv/goods_srv/global"
	"new_shop_srv/goods_srv/model"
	"new_shop_srv/goods_srv/proto"
)

//商品分类
func (s *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResp, error) {
	var categorys []model.Category
	var categorys2 []model.Category
	// TODO 不依赖gorm自己完成sql的查询和组装
	// 有外键的情况下可以使用preload预加载子分类
	res := global.DB.Where("level=?", 1).Preload("SubCategory.SubCategory").Find(&categorys)
	global.DB.Where("level=?", 1).Find(&categorys2)
	RespList := make([]*proto.CategoryInfoResp, 2)
	for _, v := range categorys2 {
		RespList = append(RespList, &proto.CategoryInfoResp{
			Id:    v.ID,
			Name:  v.Name,
			Level: v.Level,
			IsTab: v.IsTab,
		})
	}

	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResp{
		Total:    int32(res.RowsAffected),
		Data:     RespList,
		JsonData: string(b),
	}, nil
}

//获取子分类
func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListReq) (*proto.SubCategoryListResp, error) {
	var categorys model.Category

	if req.Level == 3 {
		return nil, status.Error(codes.InvalidArgument, "目录无子分类")
	}
	global.DB.Where("id=?", req.Id).Preload("SubCategory").First(&categorys)
	info := proto.CategoryInfoResp{
		Id:             categorys.ID,
		Name:           categorys.Name,
		ParentCategory: categorys.ParentCategoryId,
		Level:          categorys.Level,
		IsTab:          categorys.IsTab,
	}
	SubCategoryList := make([]*proto.CategoryInfoResp, 2)
	for _, c := range categorys.SubCategory {
		SubCategoryList = append(SubCategoryList, &proto.CategoryInfoResp{
			Id:             c.ID,
			Name:           c.Name,
			ParentCategory: c.ParentCategoryId,
			Level:          c.Level,
			IsTab:          c.IsTab,
		})
	}
	total := len(SubCategoryList)
	return &proto.SubCategoryListResp{
		Total:        int32(total),
		Info:         &info,
		SubCategorys: SubCategoryList,
	}, nil

}

func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoReq) (*proto.CategoryInfoResp, error) {
	var category model.Category
	category.Name = req.Name
	category.ParentCategoryId = req.ParentCategory
	category.Level = req.Level
	category.IsTab = req.IsTab
	global.DB.Create(&category)
	// 查询是否完成insert操作
	res := global.DB.First(&category, req.Name)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.DataLoss, "创建失败")
	}
	return &proto.CategoryInfoResp{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryId,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}, nil
}
func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryReq) (*emptypb.Empty, error) {
	var categorys []model.Category
	res := global.DB.Where("parent_category_id=?", req.Id).Find(&categorys)
	if res.RowsAffected != 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "下面还有子目录")
	}
	global.DB.Delete(model.Category{}, req.Id)
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoReq) (*emptypb.Empty, error) {
	var category model.Category
	global.DB.Model(&category).Where("id=?", req.Id).Updates(model.Category{

		Name:  req.Name,
		Level: req.Level,
		IsTab: req.IsTab,
	})
	return &emptypb.Empty{}, nil
}
