package usersModel

import (
	"GO-Store/Utils/conversion"
)

// WxAuthorizationReceiveStruct 微信登入接口数据结构体
type WxAuthorizationReceiveStruct struct {
	AvatarUrl string `json:"avatarUrl" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Gender    string `json:"gender" `
	NickName  string `json:"nickName" binding:"required"`
}

// RegisterReceiveStruct 用户注册
type RegisterReceiveStruct struct {
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

//LoginReceiveStruct 用户登入
type LoginReceiveStruct struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

//SendEmailVerCodeReceiveStruct 邮箱验证码
type SendEmailVerCodeReceiveStruct struct {
	Email string `json:"email" binding:"required,email"`
}

// ForgetReceiveStruct 用户找回密码
type ForgetReceiveStruct struct {
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

/*上面接口结构体---------------------------------------------------------------下面返回数据结构体*/

type UserInfoResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Photo    string `json:"photo"`
	Token    string `json:"token"`
}

//UserInfoResponse  生成返回用用户信息结构体
func (us User) UserInfoResponse(token string) UserInfoResponse {
	//判断用户是否为微信用户进行图片处理
	var photo string
	if len(us.Openid) <= 0 {
		photo = conversion.FormattingSrc(us.Photo)
	} else {
		photo = us.Photo
	}
	return UserInfoResponse{
		ID:       us.ID,
		UserName: us.Username,
		Photo:    photo,
		Token:    token,
	}
}
