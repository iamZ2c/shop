package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"new_shop_srv/inventory_srv/global"
	"new_shop_srv/inventory_srv/model"
	"new_shop_srv/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

//SetInv(context.Context, *InvInfo) (*emptypb.Empty, error)
//InvDetail(context.Context, *InvInfo) (*InvInfo, error)
//Sell(context.Context, *SellInfo) (*emptypb.Empty, error)
//ReBack(context.Context, *SellInfo) (*emptypb.Empty, error)

func (*InventoryServer) SetInv(ctx context.Context, req *proto.InvInfo) (*emptypb.Empty, error) {
	// 首次插入库存，二次更新库存
	inv := model.Inventory{}
	global.DB.Where("goods_id=?", req.GoodsId).First(&inv)
	fmt.Println(inv)
	inv.GoodsId = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}
func (*InventoryServer) InvDetail(ctx context.Context, req *proto.InvInfo) (*proto.InvInfo, error) {
	// 首次插入库存，二次更新库存
	inv := model.Inventory{}
	res := global.DB.First(&inv).Where("goods_id=?", req.GoodsId)
	if res.RowsAffected != 0 {
		return &proto.InvInfo{
			GoodsId: inv.GoodsId,
			Num:     inv.Stocks,
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
}
func (*InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	for _, goodsInfo := range req.GoodsInfo {

		inv := model.Inventory{}
		res := global.DB.First(&inv).Where("good_id=?", goodsInfo.GoodsId)
		if res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		if inv.Stocks < goodsInfo.Num {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存不足")
		}
		inv.Stocks -= goodsInfo.Num
		tx.Save(&inv)
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) ReBack(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	for _, goodsInfo := range req.GoodsInfo {

		inv := model.Inventory{}
		res := global.DB.Where("goods_id=?", goodsInfo.GoodsId).First(&inv)
		if res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		inv.Stocks += goodsInfo.Num
		tx.Save(&inv)
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
