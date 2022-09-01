package users

import (
	"GO-Store/Controllers"
	"GO-Store/Logic/usersLogic"
	"GO-Store/Models/users"
	"GO-Store/Utils/response"
	"GO-Store/Utils/validator"
	"github.com/gin-gonic/gin"
)

type LoginControllers struct {
	Controllers.BaseControllers
}

//WxAuthorization 微信快捷登入
func (lg LoginControllers) WxAuthorization(ctx *gin.Context) {
	WxAuthorizationReceive := new(users.WxAuthorizationReceiveStruct)
	if err := ctx.ShouldBind(WxAuthorizationReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.WxAuthorization(WxAuthorizationReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Login 登入
func (lg LoginControllers) Login(ctx *gin.Context) {
	LoginReceive := new(users.LoginReceiveStruct)
	if err := ctx.ShouldBind(LoginReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.Login(LoginReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Register 注册
func (lg LoginControllers) Register(ctx *gin.Context) {
	RegisterReceive := new(users.RegisterReceiveStruct)
	if err := ctx.ShouldBind(RegisterReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.Register(RegisterReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerCode 获取验证码(注册)
func (lg LoginControllers) SendEmailVerCode(ctx *gin.Context) {
	SendEmailVerCodeReceive := new(users.SendEmailVerCodeReceiveStruct)
	if err := ctx.ShouldBind(SendEmailVerCodeReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.SendEmailVerCode(SendEmailVerCodeReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerCodeByForget 获取邮箱验证码(忘记密码)
func (lg LoginControllers) SendEmailVerCodeByForget(ctx *gin.Context) {
	SendEmailVerCodeReceive := new(users.SendEmailVerCodeReceiveStruct)
	if err := ctx.ShouldBind(SendEmailVerCodeReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.SendEmailVerCodeByForget(SendEmailVerCodeReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Forget 找回密码
func (lg LoginControllers) Forget(ctx *gin.Context) {

	ForgetReceive := new(users.ForgetReceiveStruct)
	if err := ctx.ShouldBind(ForgetReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := usersLogic.Forget(ForgetReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
