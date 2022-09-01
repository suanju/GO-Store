package spec

// Spec 商品规格表结构
type Spec struct {
	ID      uint   `json:"id" gorm:"id"`
	GoodsId uint   `json:"goods_id" gorm:"goods_id"`
	Name    string `json:"name" gorm:"name"`
}

type List []Spec

func (Spec) TableName() string {
	return "ml_goods_spec"
}
