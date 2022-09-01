package users

import (
	"GO-Store/Logic/usersLogic"
	"GO-Store/Models/users/address"
	"GO-Store/Utils/response"
	validatorZh "GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//GetAddressTable 获取地址数据
func (us UserControllers) GetAddressTable(ctx *gin.Context) {

	addressTableReceive := new(address.RegionReceiveStruct)
	err := ctx.ShouldBind(addressTableReceive)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			msg, _ := validatorZh.ValidTrans.T(fieldError.Tag(), fieldError.Field(), fieldError.Param())
			response.Error(ctx, msg)
			return
		}
		return
	}
	results, err := usersLogic.GetAddressTable(addressTableReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
