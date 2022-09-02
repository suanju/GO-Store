package good

import (
	"GO-Store/Logic/good"
	"GO-Store/Models/goods/category"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

//GetGoodsCategory 获取商品分类
func (gd GoodsControllers) GetGoodsCategory(ctx *gin.Context) {

	GetGoodsCategoryReceive := new(category.GetGoodsCategoryStruct)
	if err := ctx.ShouldBind(GetGoodsCategoryReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := good.GetGoodsCategory(GetGoodsCategoryReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}
