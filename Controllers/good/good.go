package good

import (
	"GO-Store/Controllers"
	"GO-Store/Logic/goodLogic"
	"GO-Store/Models/goodsModel/goods"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

type GoodsControllers struct {
	Controllers.BaseControllers
}

//GetClassifyByLevel 根据分类获取商品列表
func (gd GoodsControllers) GetClassifyByLevel(ctx *gin.Context) {

	GetClassifyByLevelReceive := new(goods.GetClassifyByLevelStruct)
	if err := ctx.ShouldBind(GetClassifyByLevelReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := goodLogic.GetClassifyByLevel(GetClassifyByLevelReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}

//GetGoodInfo 获取商品详情
func (gd GoodsControllers) GetGoodInfo(ctx *gin.Context) {

	GetGoodInfoReceive := new(goods.GetGoodInfoStruct)
	if err := ctx.ShouldBind(GetGoodInfoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := goodLogic.GetGoodInfo(GetGoodInfoReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}