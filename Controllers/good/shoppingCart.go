package good

import (
	"GO-Store/Logic/goodLogic"
	"GO-Store/Models/usersModel/shoppingCart"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

//AddShoppingCart 商品添加购物车
func (gd GoodsControllers) AddShoppingCart(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	AddShoppingCartReceive := new(shoppingCart.AddShoppingCart)
	if err := ctx.ShouldBind(AddShoppingCartReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := goodLogic.AddShoppingCart(AddShoppingCartReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}

//GetShoppingCart 商品添加购物车
func (gd GoodsControllers) GetShoppingCart(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	GetShoppingCartReceive := new(shoppingCart.GetShoppingCart)
	if err := ctx.ShouldBind(GetShoppingCartReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := goodLogic.GetShoppingCart(GetShoppingCartReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)

}
