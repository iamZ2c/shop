package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_api/good_web/forms"
	"mxshop_api/good_web/global"
	"mxshop_api/good_web/proto"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	resp, err := global.GoodsSrvClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_LIST]")
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(resp.JsonData), &data)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_LIST]")
	}
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

func Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	CategoryId, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_DETAIL]")
	}
	resp, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListReq{Id: int32(CategoryId)})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_DETAIL]")
	}
	data := make(map[string]interface{}, 2)
	subCategorys := make([]map[string]interface{}, 2, 10)
	data["total"] = resp.Total
	data["id"] = resp.Info.Id
	data["name"] = resp.Info.Name
	data["level"] = resp.Info.Level
	data["parent_id"] = resp.Info.ParentCategory
	data["is_tab"] = resp.Info.IsTab

	for _, v := range resp.SubCategorys {
		c := make(map[string]interface{})
		c["id"] = v.Id
		c["name"] = v.Name
		c["level"] = v.Level
		c["parent_id"] = v.ParentCategory
		subCategorys = append(subCategorys, c)
	}
	data["sub_category_lit"] = subCategorys
	ctx.JSON(
		http.StatusOK,
		data,
	)
	return
}

func New(ctx *gin.Context) {
	categoryForm := forms.CategoryForm{}
	err := ctx.ShouldBindUri(&categoryForm)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_NEW]")
	}
	resp, err := global.GoodsSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoReq{
		Name:           categoryForm.Name,
		ParentCategory: categoryForm.ParentCategory,
		Level:          categoryForm.Level,
		IsTab:          categoryForm.IsTab,
	})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_NEW]")
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message":     "success",
			"category_id": resp.Id,
		},
	)
	return
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_DELETE]")
	}

	_, err = global.GoodsSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryReq{Id: int32(categoryId)})
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_DELETE]")
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)

}
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_DELETE]")
	}
	UpdatecategoryForm := forms.UpdateCategoryForm{}
	err = ctx.ShouldBindUri(&UpdatecategoryForm)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_UPDATE]")
	}
	req := &proto.CategoryInfoReq{}
	if UpdatecategoryForm.IsTab != nil {
		req = &proto.CategoryInfoReq{
			Id:   int32(categoryId),
			Name: UpdatecategoryForm.Name,
		}
	} else {
		req = &proto.CategoryInfoReq{
			Id:    int32(categoryId),
			Name:  UpdatecategoryForm.Name,
			IsTab: *UpdatecategoryForm.IsTab,
		}
	}
	_, err = global.GoodsSrvClient.UpdateCategory(context.Background(), req)
	if err != nil {
		global.Check(ctx, err, "[CATEGORY_UPDATE]")
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
		},
	)
}
