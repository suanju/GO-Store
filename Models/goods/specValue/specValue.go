package spec

// Value 商品规格值结构
type Value struct {
	ID      uint   `json:"id" gorm:"id"`
	GoodsId uint   `json:"goods_id" gorm:"goods_id"`
	SpecId  uint   `json:"spec_id" gorm:"spec_id"`
	Value   string `json:"value" gorm:"value"`
}
type ValueList []Value

func (Value) TableName() string {
	return "ml_goods_spec_value"
}
