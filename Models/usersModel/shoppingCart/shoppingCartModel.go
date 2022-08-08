package shoppingCart

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/goodsModel/goods"
	"GO-Store/Models/goodsModel/item"
	"time"
)

//ShoppingCart 用户购物车表
type ShoppingCart struct {
	ID        uint      `json:"id" gorm:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UserID    uint      `json:"user_id" gorm:"user_id"`
	GoodsId   uint      `json:"goods_id" gorm:"goods_id"`
	ItemID    uint      `json:"item_id" gorm:"item_id"`
	Number    int64     `json:"number" gorm:"Number"`

	//关联模型
	GoodsInfo goods.Goods `json:"goods_info" gorm:"foreignKey:GoodsId"`
	Item      item.Item   `json:"item" gorm:"foreignKey:ItemID"`
}
type List []ShoppingCart

func (ShoppingCart) TableName() string {
	return "ml_user_shopping_cart"
}

//AddGood 添加购物车
func (sc *ShoppingCart) AddGood() bool {

	err := Mysql.Db.Create(&sc).Error
	if err != nil {
		return false
	}
	return true
}

func (ls *List) GetList(uid uint) error {
	err := Mysql.Db.Where(ShoppingCart{UserID: uid}).Preload("GoodsInfo").Preload("Item").Find(&ls).Error
	return err
}
