package users

import (
	"GO-Store/Controllers/users"
	"GO-Store/Middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	usersRouter := Router.Group("user").Use(Middlewares.VerificationToken())
	{
		userControllers := new(users.UserControllers)
		usersRouter.POST("/getAddressTable", userControllers.GetAddressTable)
		usersRouter.POST("/setAddress", userControllers.SetAddress)
		usersRouter.POST("/getAddressList", userControllers.GetAddressList)
	}
}
