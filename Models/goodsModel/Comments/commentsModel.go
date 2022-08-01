package Comments

import "GO-Store/Models/common"

type Comments struct {
	common.PublicModel
	GoodsId     uint   `json:"goods_id" gorm:"goods_id"`
	UserID      uint   `json:"user_id" gorm:"user_id"`
	SpecValueId string `json:"spec_value_id" gorm:"spec_value_id"`
	Text        string `json:"text" gorm:"text"`
	Image       string `json:"image" gorm:"image"`
}

func (Comments) TableName() string {
	return "ml_goods_comments"
}
