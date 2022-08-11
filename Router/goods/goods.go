package goods

import (
	"GO-Store/Controllers/good"
	"GO-Store/Middlewares"
	"github.com/gin-gonic/gin"
)

type GoodRouter struct {
}

func (g *GoodRouter) InitRouter(Router *gin.RouterGroup) {
	goodRouter := Router.Group("good").Use()
	{
		goodControllers := new(good.GoodsControllers)
		goodRouter.POST("/getGoodsCategory", goodControllers.GetGoodsCategory)
		goodRouter.POST("/getClassifyByLevel", goodControllers.GetClassifyByLevel)
		goodRouter.POST("/getGoodInfo", Middlewares.VerificationToken(), goodControllers.GetGoodInfo)
		goodRouter.POST("/collection", Middlewares.VerificationToken(), goodControllers.Collection)
		goodRouter.POST("/addShoppingCart", Middlewares.VerificationToken(), goodControllers.AddShoppingCart)
		goodRouter.POST("/getShoppingCart", Middlewares.VerificationToken(), goodControllers.GetShoppingCart)
		goodRouter.POST("/delShoppingCart", Middlewares.VerificationToken(), goodControllers.DelShoppingCart)
		goodRouter.POST("/modifyInventory", Middlewares.VerificationToken(), goodControllers.ModifyInventory)
	}
}
