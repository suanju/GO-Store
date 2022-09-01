package home

import (
	"GO-Store/Controllers"
	"GO-Store/Logic/homeLogic"
	"GO-Store/Models/home/banner"
	"GO-Store/Models/home/special"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

type HomesControllers struct {
	Controllers.BaseControllers
}

//GetBannerList 获取首页轮播图
func (hm HomesControllers) GetBannerList(ctx *gin.Context) {

	GetBannerListReceive := new(banner.GetBannerListStruct)
	if err := ctx.ShouldBind(GetBannerListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := homeLogic.GetBannerList(GetBannerListReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}

//GetSpecialList 获取首页会场信息
func (hm HomesControllers) GetSpecialList(ctx *gin.Context) {

	GetSpecialListReceive := new(special.GetSpecialListStruct)
	if err := ctx.ShouldBind(GetSpecialListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := homeLogic.GetSpecialList(GetSpecialListReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}
