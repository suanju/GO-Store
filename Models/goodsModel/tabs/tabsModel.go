package tabs

import (
	"GO-Store/Databases/Mysql"
	"fmt"
)

// Tabs 商品规格值结构
type Tabs struct {
	ID              uint   `json:"id" gorm:"id"`
	Name            string `json:"name" gorm:"name"`
	AssociatedClass string `json:"	associated_class" gorm:"	associated_class"`
}
type TabList []Tabs

func (Tabs) TableName() string {
	return "ml_goods_tabs"
}

func (l *TabList) GetAllList() (err error) {
	err = Mysql.Db.Find(l).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	return nil
}

func (t *Tabs) GetInfoById() (err error) {
	err = Mysql.Db.Find(t).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	return nil
}
