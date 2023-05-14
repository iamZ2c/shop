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

		// redis分布式锁
		Mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%v", goodsInfo.GoodsId))
		// 加锁
		if err := Mutex.Lock(); err != nil {
			return nil, status.Error(codes.Internal, "锁获取异常")
		}
		// 释放锁
		if ok, err := Mutex.Unlock(); !ok || err != nil {
			return nil, status.Error(codes.Internal, "释放锁异常 ")
		}

		// 使用mysql for update 悲观锁
		//res := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("good_id=?", goodsInfo.GoodsId).First(&inv)

		for true {
			res := tx.Where("good_id=?", goodsInfo.GoodsId).First(&inv)
			if res.RowsAffected == 0 {
				tx.Rollback()
				return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
			}
			if inv.Stocks < goodsInfo.Num {
				tx.Rollback()
				return nil, status.Error(codes.InvalidArgument, "库存不足")
			}
			inv.Stocks -= goodsInfo.Num
			// 乐观锁，更新的时候查看版本呢是否与前面一致
			res = tx.Select("stocks", "version").Where("goods_id=? AND version=?", goodsInfo.GoodsId, goodsInfo.Version)
			if res.RowsAffected != 0 {
				break
			}
		}
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
