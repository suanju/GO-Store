package goods

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/common"
	"GO-Store/Models/goods/comments"
	"GO-Store/Models/goods/item"
	"GO-Store/Models/goods/service"
	"GO-Store/Models/goods/spec"
	spec2 "GO-Store/Models/goods/specValue"
	"GO-Store/Models/users/likeGoods"
	"fmt"
	"gorm.io/gorm"
)

// Goods 商品分类表结构
type Goods struct {
	common.PublicModel
	Type              int8    `json:"type" gorm:"type"`
	Name              string  `json:"name" gorm:"name"`
	Code              string  `json:"code" gorm:"code"`
	FirstCateId       int64   `json:"first_cate_id" gorm:"first_cate_id"`
	SecondCateId      int64   `json:"second_cate_id" gorm:"second_cate_id"`
	ThirdCateId       int64   `json:"third_cate_id" gorm:"third_cate_id"`
	BranId            int64   `json:"brand_id" gorm:"brand_id"`
	UnitId            int64   `json:"unit_id" gorm:"unit_id"`
	Status            int8    `json:"status" gorm:"status"`
	Image             string  `json:"image" gorm:"image"`
	ShufflingImg      string  `json:"shuffling_img" gorm:"shuffling_img"`
	Video             string  `json:"video" gorm:"video"`
	Remark            string  `json:"remark" gorm:"remark"`
	Content           string  `json:"content" gorm:"content"`
	Sort              int64   `json:"sort" gorm:"sort"`
	SalesActual       int64   `json:"sales_actual" gorm:"sales_actual"`
	Clicks            int64   `json:"clicks" gorm:"clicks"`
	Tabs              string  `json:"tabs" gorm:"tabs"`
	SpecType          int8    `json:"spec_type" gorm:"spec_type"`
	Price             float64 `json:"price" gorm:"price"`
	Stock             int64   `json:"stock" gorm:"stock"`
	ExpressType       int8    `json:"express_type" gorm:"express_type"`
	ExpressMoney      float64 `json:"express_money" gorm:"express_money"`
	ExpressTemplateId int64   `json:"express_template_id" gorm:"express_template_id"`
	Del               int8    `del:"name" gorm:"del"`
	StockWarn         int64   `json:"stock_warn" gorm:"stock_warn"`
	ColumnIds         string  `json:"column_ids" gorm:"column_ids"`
	SalesVirtual      int64   `json:"sales_virtual" gorm:"sales_virtual"`
	SortWeight        int64   `json:"sort_weight" gorm:"sort_weight"`
	IsShowStock       int8    `json:"is_show_stock" gorm:"is_show_stock"`
	IsMember          int8    `json:"is_member" gorm:"is_member"`
	DeliveryType      string  `json:"delivery_type" gorm:"delivery_type"`
	AfterPay          string  `json:"after_pay" gorm:"after_pay"`
	AfterDelivery     int8    `json:"after_delivery" gorm:"after_delivery"`
	DeliveryContent   int8    `json:"delivery_content" gorm:"delivery_content"`

	//商品表拓展
	Spec      spec.List             `json:"spec" gorm:"foreignKey:GoodsId"`
	SpecValue spec2.ValueList       `json:"spec_value" gorm:"foreignKey:GoodsId"`
	Item      item.List             `json:"item" gorm:"foreignKey:GoodsId"`
	SeverList service.ServicesLists `json:"sever_list" gorm:"foreignKey:GoodsId"`
	Comments  comments.List         `json:"comments" gorm:"foreignKey:GoodsId"`
	//用户与商品之间的关系
	Like likeGoods.LikeGoods `json:"like" gorm:"foreignKey:GoodsId"`
}
type GoodList []Goods

func (Goods) TableName() string {
	return "ml_goods"
}

func (g *Goods) SelectById(id int64, uid uint) error {
	err := Mysql.Db.Preload("SpecValue").Preload("Spec").Preload("Item").Preload("SeverList.Service").Preload("Comments.UserInfo").Preload("Like", " user_id  =  ?", uid).Find(&g, id).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	if g.ID <= 0 {
		return fmt.Errorf("商品不存在")
	}
	return nil
}

//SelectByLevel 根据分类查询商品
func (l *GoodList) SelectByLevel(level int64, pid int64, info common.PageInfo) error {
	switch level {
	case 1:
		err := Mysql.Db.Where(&Goods{FirstCateId: pid}).Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize).Find(&l).Error
		if err != nil {
			return fmt.Errorf("查询失败")
		}
		break
	case 2:
		err := Mysql.Db.Where(&Goods{SecondCateId: pid}).Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize).Find(&l).Error
		if err != nil {
			return fmt.Errorf("查询失败")
		}
		break
	case 3:
		err := Mysql.Db.Where(&Goods{ThirdCateId: pid}).Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize).Find(&l).Error
		if err != nil {
			return fmt.Errorf("查询失败")
		}
		break
	default:
		return fmt.Errorf("等级错误")
	}
	return nil
}

//SelectByTab 根据TABS查询商品
func (l *GoodList) SelectByTab(data *GetTabsInfo, tabId uint, page int, pageSize int) error {
	err := Mysql.Db.Raw("SELECT * FROM `ml_goods` WHERE find_in_set(?,`tabs`) LIMIT ? OFFSET ?", tabId, pageSize, (page-1)*pageSize).Scan(&data.List).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	err = Mysql.Db.Raw("SELECT COUNT(*) FROM `ml_goods` WHERE find_in_set(?,`tabs`)", tabId).Scan(&data.Size).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	return nil
}

//ClicksAdd 点击量增加
func (g *Goods) ClicksAdd() {
	Mysql.Db.Model(&g).Updates(map[string]interface{}{"clicks": gorm.Expr("Clicks  + ?", 1)})
}
