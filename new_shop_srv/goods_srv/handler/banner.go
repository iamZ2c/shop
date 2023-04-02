package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"new_shop_srv/goods_srv/global"
	"new_shop_srv/goods_srv/model"
	"new_shop_srv/goods_srv/proto"
)

func (s *GoodsServer) BannerList(context.Context, *emptypb.Empty) (*proto.BannerListResp, error) {
	var banners []model.Banner
	global.DB.Find(&banners)
	var c int64
	global.DB.Find(&banners).Count(&c)
	data := make([]*proto.BannerResp, 2)
	for _, value := range banners {
		data = append(data, &proto.BannerResp{
			Id:    value.ID,
			Image: value.Image,
			Url:   value.Url,
			Index: value.Index,
		})
	}
	return &proto.BannerListResp{
		Total: int32(c),
		Data:  data,
	}, nil

}

func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerReq) (*proto.BannerResp, error) {
	var banner model.Banner
	banner.ID = req.Id
	banner.Url = req.Url
	banner.Image = req.Image
	banner.Index = req.Index
	res := global.DB.Create(&banner)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, res.Error.Error())
		}
	}
	return &proto.BannerResp{
		Id:    banner.ID,
		Index: banner.Index,
		Url:   banner.Url,
		Image: banner.Image,
	}, nil

}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerReq) (*emptypb.Empty, error) {
	res := global.DB.First(&model.Banner{}, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "广告不存在")
	}
	global.DB.Delete(req.Id)
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerReq) (*emptypb.Empty, error) {
	var banner model.Banner
	res := global.DB.First(&banner, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "广告不存在")
	}
	banner.Url = req.Url
	banner.Image = req.Image
	banner.Index = req.Index
	global.DB.Save(&banner)
	return &emptypb.Empty{}, nil
}
