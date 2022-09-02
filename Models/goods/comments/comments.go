package comments

import (
	"GO-Store/Models/common"
	"GO-Store/Models/users"
)

type Comments struct {
	common.PublicModel
	GoodsId     uint   `json:"goods_id" gorm:"goods_id"`
	UserID      uint   `json:"user_id" gorm:"user_id"`
	SpecValueId string `json:"spec_value_id" gorm:"spec_value_id"`
	Text        string `json:"text" gorm:"text"`
	Image       string `json:"image" gorm:"image"`

	UserInfo users.User `json:"user_info" gorm:"foreignKey:UserID"`
}

type List []Comments

func (Comments) TableName() string {
	return "ml_goods_comments"
}
