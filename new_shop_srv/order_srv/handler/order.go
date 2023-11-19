package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"math/rand"
	"new_shop_srv/order_srv/global"
	"new_shop_srv/order_srv/model"
	"new_shop_srv/order_srv/proto"
	"time"
)

type OrderServer struct {
	proto.UnimplementedOrderServer
}

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
	// TODO 使用save后，是否会给shopCart的id传值，（应该会）,已测试，会。
	res = global.DB.Save(&shopCart)
	if res.Error != nil {
		return nil, status.Errorf(codes.Unimplemented, "save error")
	}
	return &proto.ShopCartInfoResp{
		Id: shopCart.ID,
	}, nil
}

//
func (*OrderServer) UpdateCartItem(ctx context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	var shopCart model.ShoppingCart
	if res := global.DB.First(req.Id).Find(&shopCart); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "shop cart not found")
	} else {
		if req.Nums > 0 {
			shopCart.Nums += req.Nums
		}
		shopCart.Checked = req.Checked
		shopCart.User = req.UserId
		res = global.DB.Save(&shopCart)
		if res.Error != nil {
			return nil, status.Errorf(codes.Unimplemented, "save error")
		}
		return &emptypb.Empty{}, nil
	}
}

func (*OrderServer) DeleteCartItem(ctx context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	if res := global.DB.Delete(req.Id).Find(&model.ShoppingCart{}); res.RowsAffected != 0 {
		return &emptypb.Empty{}, nil
	} else {
		return nil, status.Errorf(codes.NotFound, "shop cart not found")
	}
}

//rpc Create(OrderReq) returns(OrderInfoResp); // 创建订单
//
//rpc OrderList(OrderFilterReq)returns(OrderListResp); // 订单列表
//
//rpc OrderDetail(OrderReq) returns(OrderInfoDetailResp);//
//
//rpc UpdateOrderStatus(OrderStatus) returns(google.protobuf.Empty); //修改订单状态

func (*OrderServer) OrderList(ctx context.Context, req *proto.OrderFilterReq) (*proto.OrderListResp, error) {
	var orders []model.OrderInfo
	var resp *proto.OrderListResp
	var total int64
	global.DB.Where(&model.OrderInfo{User: req.UserId}).Count(&total)
	resp.Total = int32(total)

	global.DB.Find(&orders).Scopes(Paginate(int(req.Pages), int(req.PagePerNums)))
	ordersInfoList := make([]*proto.OrderInfoResp, 2)
	for _, orderInfo := range orders {
		order := proto.OrderInfoResp{
			Id:      orderInfo.ID,
			UserId:  orderInfo.User,
			OrderSn: orderInfo.OrderSn,
			PayType: orderInfo.PayType,
			Status:  orderInfo.Status,
			Post:    orderInfo.Post,
			Total:   orderInfo.Total,
			Address: orderInfo.Address,
			Name:    orderInfo.SignerName,
			Mobile:  orderInfo.SingerMobile,
		}
		ordersInfoList = append(ordersInfoList, &order)
	}
	resp.Data = ordersInfoList
	return resp, nil
}

func (*OrderServer) OrderDetail(ctx context.Context, req *proto.OrderReq) (*proto.OrderInfoDetailResp, error) {
	var orderInfo model.OrderInfo
	var resp proto.OrderInfoDetailResp
	res := global.DB.Where(&model.OrderInfo{BaseModel: model.BaseModel{ID: req.Id}, User: req.UserId}).Find(&orderInfo)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}
	resp.OrderInfo = &proto.OrderInfoResp{
		Id:      orderInfo.ID,
		UserId:  orderInfo.User,
		OrderSn: orderInfo.OrderSn,
		PayType: orderInfo.PayType,
		Status:  orderInfo.Status,
		Post:    orderInfo.Post,
		Total:   orderInfo.Total,
		Address: orderInfo.Address,
		Name:    orderInfo.SignerName,
		Mobile:  orderInfo.SingerMobile,
	}

	var orderGoods []model.OrderGoods
	global.DB.Where(&model.OrderGoods{Order: req.Id}).Find(&orderGoods)

	orderItems := make([]*proto.OrderItemResp, 2)
	for _, item := range orderGoods {
		oir := proto.OrderItemResp{
			Id:         item.ID,
			OrderId:    item.Order,
			GoodsId:    item.Goods,
			GoodsName:  item.GoodsName,
			GoodsImage: item.GoodsImage,
			GoodsPrice: int32(item.GoodsPrice),
		}
		orderItems = append(orderItems, &oir)
	}
	resp.Goods = orderItems
	return &resp, nil
}

func (*OrderServer) Create(ctx context.Context, req *proto.OrderReq) (*proto.OrderInfoResp, error) {
	var goodsIds []int32
	var shopCarts []model.ShoppingCart
	var idAndNums map[int32]int32
	if res := global.DB.Where(&model.ShoppingCart{User: req.UserId, Checked: true}).Find(&shopCarts); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "未选择商品")
	}
	for _, shopcart := range shopCarts {
		goodsIds = append(goodsIds, shopcart.Goods)
		idAndNums[shopcart.Goods] = shopcart.Nums
	}

	// 查询 goodssrv 拿到价格,并且计算总价
	goodsResp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodsIds})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "批量获取商品失败")
	}

	var orderTotalPrice float32
	var orderGoods []*model.OrderGoods
	var goodsInfo []*proto.InvInfo
	for _, goods := range goodsResp.Data {

		orderTotalPrice += goods.ShopPrice * float32(idAndNums[goods.Id])
		// 先生成好，后续不用从新生成插入表
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      goods.Id,
			GoodsName:  goods.Name,
			GoodsImage: goods.GoodsFrontImage,
			GoodsPrice: goods.ShopPrice,
			Nums:       idAndNums[goods.Id],
		})

		goodsInfo = append(goodsInfo, &proto.InvInfo{
			GoodsId: goods.Id,
			Num:     idAndNums[goods.Id],
		})
	}
	// 库存服务库存扣减
	_, err = global.InventoryClient.Sell(context.Background(), &proto.SellInfo{GoodsInfo: goodsInfo})
	if err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "扣减库存失败")
	}
	// 存入订单数据
	orderInfo := model.OrderInfo{
		User:         req.UserId,
		OrderSn:      GenerateOrderSn(req.UserId),
		Address:      req.Address,
		SignerName:   req.Name,
		SingerMobile: req.Mobile,
		Post:         req.Post,
		OrderMount:   orderTotalPrice,
	}
	tx := global.DB.Begin()
	res := tx.Save(&orderInfo)
	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "插入订单失败")
	}
	// 存入订单物品关联数据,前面生成过了
	for _, goods := range orderGoods {
		goods.Order = orderInfo.ID
	}
	// 插入数据
	res = tx.CreateInBatches(orderGoods, 100)
	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "插入订单失败")
	}
	// 删除购物车得商品
	tx.Where(&model.ShoppingCart{User: req.UserId, Checked: true}).Delete(&model.ShoppingCart{})
	if res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "插入订单失败")
	}
	return &proto.OrderInfoResp{
		Id:      orderInfo.ID,
		OrderSn: orderInfo.OrderSn,
		Total:   orderInfo.OrderMount,
	}, nil
}

func GenerateOrderSn(UserId int32) string {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		UserId, rand.Intn(90)+10,
	)
	return orderSn
	//	1-15
}
