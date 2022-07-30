package Router

import (
	"GO-Store/Middlewares"
	"GO-Store/Router/goods"
	"GO-Store/Router/home"
	usersRouter "GO-Store/Router/users"
	"github.com/gin-gonic/gin"
)

type RoutersGroup struct {
	Users usersRouter.RouterGroup
	Home  home.RouterGroup
	Good  goods.GoodRouter
}

var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	router.Use(Middlewares.Cors())
	PrivateGroup := router.Group("")
	PrivateGroup.Use()
	{
		//静态资源访问
		router.Static("/assets", "./Assets/")
		RoutersGroupApp.Users.LoginRouter.InitLoginRouter(PrivateGroup)
		RoutersGroupApp.Users.InitRouter(PrivateGroup)
		RoutersGroupApp.Home.MainRouter.InitRouter(PrivateGroup)
		RoutersGroupApp.Good.InitRouter(PrivateGroup)
	}

	router.Run()
}
