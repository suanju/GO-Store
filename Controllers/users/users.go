package users

import (
	"GO-Store/Controllers"
	"GO-Store/Logic/users"
	"GO-Store/Models/users/address"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
	base Controllers.BaseControllers
}

//SetAddress 设置用户地址
func (us UserControllers) SetAddress(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	setAddressReceive := new(address.SetAddressStruct)
	if err := ctx.ShouldBind(setAddressReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SetAddress(userID, setAddressReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetAddressList 获取用户地址列表
func (us UserControllers) GetAddressList(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	results, err := users.GetAddressList(userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
