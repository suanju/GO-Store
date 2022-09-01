package category

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/common"
)

// Category 商品分类表结构
type Category struct {
	common.PublicModel
	Name    string `json:"name" gorm:"name"`
	Pid     int64  `json:"pid" gorm:"pid"`
	Level   int64  `json:"level" gorm:"level"`
	Sort    int64  `json:"sort" gorm:"sort"`
	IsShow  int8   `json:"is_show" gorm:"is_show"`
	Picture string `json:"picture" gorm:"picture"`
	Remark  string `json:"remark" gorm:"remark"`
}
type GoodsCategoryList []Category

func (Category) TableName() string {
	return "ml_goods_category"
}

//SelectAll 获取用户地址列表
func (l *GoodsCategoryList) SelectAll() bool {
	err := Mysql.Db.Find(&l).Error
	if err != nil {
		return false
	}
	return true
}
