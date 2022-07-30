package good

import (
	"GO-Store/Logic/goodLogic"
	"GO-Store/Models/goodsModel/category"
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
	results, err := goodLogic.GetGoodsCategory(GetGoodsCategoryReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}
