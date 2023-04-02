package banner

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_api/good_web/forms"
	"mxshop_api/good_web/global"
	"mxshop_api/good_web/proto"
	"net/http"
)

//CreateBanner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*BannerResp, error)
//DeleteBanner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
//UpdateBanner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
func Get(ctx *gin.Context) {
	resp, err := global.GoodsSrvClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		global.Check(ctx, err, "[BANNER_GET]")
	}
	bannerList := make([]map[string]interface{}, 2)
	for _, v := range resp.Data {
		banner := make(map[string]interface{})
		banner["id"] = v.Id
		banner["url"] = v.Url
		banner["image"] = v.Image
		banner["index"] = v.Index
		bannerList = append(bannerList, banner)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"total":   resp.Total,
				"banners": bannerList,
			},
		},
	)
	return
}

func New(ctx *gin.Context) {
	b := forms.BannerForm{}
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		global.Check(ctx, err, "[NEW_BANNER]")
	}

	resp, err := global.GoodsSrvClient.CreateBanner(context.Background(), &proto.BannerReq{

		Index: int32(b.Index),
		Image: b.Image,
		Url:   b.Url,
	})
	if err != nil {
		global.Check(ctx, err, "[NEW_BANNER]")
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"good_id": resp.Id,
			},
		},
	)
	return
}

func Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id := global.Str2Int(idStr)
	_, err := global.GoodsSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: id})
	if err != nil {
		global.Check(ctx, err, "[NEW_BANNER]")
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"good_id": id,
			},
		},
	)
}

func Update(ctx *gin.Context) {
	b := forms.BannerForm{}
	idStr := ctx.Param("id")
	id := global.Str2Int(idStr)
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		global.Check(ctx, err, "[UPDATE_BANNER]")
	}

	_, err = global.GoodsSrvClient.UpdateBanner(context.Background(), &proto.BannerReq{
		Id:    id,
		Index: int32(b.Index),
		Image: b.Image,
		Url:   b.Url,
	})
	if err != nil {
		global.Check(ctx, err, "[UPDATE_BANNER]")
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"good_id": id,
			},
		},
	)
}
