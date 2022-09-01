package item

// Item 商品SUK结构
type Item struct {
	ID           uint    `json:"id" gorm:"id"`
	Image        string  `json:"image" gorm:"image"`
	GoodsId      uint    `json:"goods_id" gorm:"goods_id"`
	SpecValueIds string  `json:"spec_value_ids" gorm:"spec_value_ids"`
	SpecValueStr string  `json:"spec_value_str" gorm:"spec_value_str"`
	Price        float64 `json:"price" gorm:"price"`
	Stock        int64   `json:"stock" gorm:"stock"`
}

type List []Item

func (Item) TableName() string {
	return "ml_goods_item"
}
