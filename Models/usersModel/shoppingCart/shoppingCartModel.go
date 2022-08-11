package shoppingCart

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/goodsModel/goods"
	"GO-Store/Models/goodsModel/item"
	"fmt"
	"gorm.io/gorm"
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

//GetList 获取购物车列表
func (ls *List) GetList(uid uint) error {
	err := Mysql.Db.Where(ShoppingCart{UserID: uid}).Preload("GoodsInfo").Preload("Item").Find(&ls).Error
	return err
}

//Empty 清空购物车
func (sc ShoppingCart) Empty(uid uint) error {
	err := Mysql.Db.Where(ShoppingCart{UserID: uid}).Delete(&sc).Error
	return err
}

//DelByID 删除根据ID
func (sc ShoppingCart) DelByID(id uint) error {
	//判断购物车数据是否存在
	err := Mysql.Db.Where(ShoppingCart{ID: id}).Find(&sc).Error
	if err != nil {
		return fmt.Errorf("数据不存在")
	}
	err = Mysql.Db.Where(ShoppingCart{ID: id}).Delete(&sc).Error
	return err
}

//ModifyInventory 修改购物车数量
func (sc *ShoppingCart) ModifyInventory(uid uint, id uint, tp string) error {
	sc.ID = id
	err := Mysql.Db.Where(ShoppingCart{ID: id}).Find(&sc).Error
	if err != nil {
		return err
	}
	if sc.UserID != uid {
		return fmt.Errorf("不可删除")
	}
	if tp == "+" {
		Mysql.Db.Model(&sc).Updates(map[string]interface{}{"number": gorm.Expr("number  + ?", 1)})
	} else {
		Mysql.Db.Model(&sc).Updates(map[string]interface{}{"number": gorm.Expr("number  - ?", 1)})
	}

	return nil
}
