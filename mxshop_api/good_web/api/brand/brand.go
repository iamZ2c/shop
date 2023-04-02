package brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/forms"
	"mxshop_api/good_web/global"
	"mxshop_api/good_web/proto"
	"net/http"
)

// CategoryBrandList 获取所有分类品牌信息
func CategoryBrandList(ctx *gin.Context) {
	resp, err := global.GoodsSrvClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterReq{})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_BRAND_LIST_SRV]")
		return
	}
	data := make([]map[string]any, 2)
	for _, v := range resp.Data {
		d := make(map[string]interface{})
		d["id"] = v.Id
		d["brand"] = map[string]any{
			"logo": v.Brand.Logo,
			"name": v.Brand.Name,
		}
		d["category"] = map[string]any{
			"id":        v.Category.Id,
			"parent_id": v.Category.ParentCategory,
			"is_tab":    v.Category.IsTab,
			"name":      v.Category.Name,
			"level":     v.Category.Level,
		}
		data = append(data, d)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"total":                resp.Total,
				"category_brands_info": data,
			},
		},
	)
}

// GetCategoryBrandList 根据分类查品牌
func GetCategoryBrandList(ctx *gin.Context) {

	// 传入的是categoryid
	id := global.Str2Int(ctx.Param("id"))
	resp, err := global.GoodsSrvClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoReq{
		Id: id,
	})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_BRAND_LIST_SRV]")
		return
	}
	data := make([]map[string]interface{}, 2)
	for _, v := range resp.Data {
		b := make(map[string]interface{})
		b["id"] = v.Id
		b["logo"] = v.Logo
		b["name"] = v.Name
		data = append(data, b)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"total":  resp.Total,
				"brands": data,
			},
		},
	)
	return
}

func Get(ctx *gin.Context) {
	resp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterReq{
		Pages:       global.Str2Int(ctx.Query("pages")),
		PagePerNums: global.Str2Int(ctx.Query("page_per_nums")),
	})
	if err != nil {
		global.Check(ctx, err, "[GET_BRAND]")
		return
	}

	data := make([]map[string]interface{}, 2)
	for _, brand := range resp.Data {
		b := make(map[string]interface{})
		b["id"] = brand.Id
		b["name"] = brand.Name
		b["logo"] = brand.Logo
		data = append(data, b)
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"brands": data,
				"total":  resp.Total,
			},
		},
	)

}

func New(ctx *gin.Context) {

	b := forms.BrandForm{}
	err := ctx.ShouldBindJSON(b)
	if err != nil {
		global.Check(ctx, err, "[CREATE_BRAND]")
		return
	}
	resp, err := global.GoodsSrvClient.CreateBrand(
		context.Background(),
		&proto.BrandReq{
			Name: b.Name,
			Logo: b.Logo,
		},
	)
	if err != nil {
		global.Check(ctx, err, "[CREATE_BRAND_SRV]")
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data": gin.H{
				"logo": resp.Logo,
				"name": resp.Name,
				"id":   resp.Id,
			},
		},
	)
}

func Delete(ctx *gin.Context) {
	id := global.Str2Int(ctx.Param("id"))
	_, err := global.GoodsSrvClient.DeleteBrand(context.Background(), &proto.BrandReq{
		Id: id,
	})
	if err != nil {
		global.Check(ctx, err, "[DELETE_BRAND]")
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
}

func Update(ctx *gin.Context) {
	id := global.Str2Int(ctx.Param("id"))
	b := forms.BrandForm{}
	err := ctx.ShouldBindJSON(b)
	if err != nil {
		global.Check(ctx, err, "[CREATE_BRAND]")
		return
	}
	_, err = global.GoodsSrvClient.UpdateBrand(context.Background(), &proto.BrandReq{
		Id:   id,
		Name: b.Name,
		Logo: b.Logo,
	})
	if err != nil {
		global.Check(ctx, err, "[CREATE_BRAND_SRV]")
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
	return
}
