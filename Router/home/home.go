package home

import (
	"GO-Store/Controllers/home"
	"github.com/gin-gonic/gin"
)

type MainRouter struct {
}

func (s *MainRouter) InitRouter(Router *gin.RouterGroup) {
	homeRouter := Router.Group("home").Use()
	{
		homeControllers := new(home.HomesControllers)
		homeRouter.POST("/getBannerList", homeControllers.GetBannerList)
		homeRouter.POST("/getSpecialList", homeControllers.GetSpecialList)
	}
}
