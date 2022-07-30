package goods

import (
	"GO-Store/Controllers/good"
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
		goodRouter.POST("/getGoodInfo", goodControllers.GetGoodInfo)
	}
}
