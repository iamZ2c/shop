package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"new_shop_srv/order_srv/global"
	"new_shop_srv/order_srv/model"
	"new_shop_srv/order_srv/proto"
)

type OrderServer struct {
	proto.UnimplementedOrderServer
}

// CartItemList 获取用户购物车列表
func (*OrderServer) CartItemList(ctx context.Context, req *proto.UserInfo) (*proto.CartItemListResp, error) {
	var shopCarts []*model.ShoppingCart
	//res := global.DB.Where("user=?", req.Id).Find(&shopCarts)
	if res := global.DB.Where(&model.ShoppingCart{User: req.Id}).Find(&shopCarts); res.Error != nil {
		return nil, status.Errorf(codes.Internal, "server error")
	} else {
		data := make([]*proto.ShopCartInfoResp, 5)
		for _, v := range shopCarts {
			data = append(data, &proto.ShopCartInfoResp{
				Id:      v.ID,
				UserId:  v.User,
				Nums:    v.Nums,
				GoodsId: v.Goods,
				Checked: v.Checked,
			})
		}
		resp := &proto.CartItemListResp{
			Total: int32(res.RowsAffected),
			Data:  data,
		}
		return resp, nil
	}
}

// CreateCartItem 商品添加到购物车
func (*OrderServer) CreateCartItem(ctx context.Context, req *proto.CartItemReq) (*proto.ShopCartInfoResp, error) {
	var shopCart model.ShoppingCart
	res := global.DB.Where(&model.ShoppingCart{User: req.UserId, Goods: req.GoodsId}).Find(&shopCart)
	if res.RowsAffected != 0 {
		shopCart.Nums += req.Nums
		shopCart.Checked = req.Checked
		shopCart.User = req.UserId
	} else {
		res = global.DB.Create(&model.ShoppingCart{
			User:    req.UserId,
			Goods:   req.GoodsId,
			Nums:    req.Nums,
			Checked: false,
		})
	}
	// TODO 使用save后，是否会给shopCart的id传值，（应该会）
	res = global.DB.Save(&shopCart)
	if res.Error != nil {
		return nil, status.Errorf(codes.Unimplemented, "save error")
	}
	return &proto.ShopCartInfoResp{
		Id: shopCart.ID,
	}, nil
}
