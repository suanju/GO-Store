package good

import (
	"GO-Store/Logic/goodLogic"
	"GO-Store/Models/goodsModel/tabs"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

//GetTabs 获取tabs信息
func (gd GoodsControllers) GetTabs(ctx *gin.Context) {
	GetTabsReceive := new(tabs.GetTabsStruct)
	if err := ctx.ShouldBind(GetTabsReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := goodLogic.GetTabs(GetTabsReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}
