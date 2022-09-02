package shoppingCart

type AddShoppingCart struct {
	GoodsID uint  `json:"goods_id"`
	ItemId  uint  `json:"item_id"`
	Number  int64 `json:"number"`
}

type GetShoppingCart struct {
}

type DelShoppingCart struct {
	ShoppingCartID uint `json:"shoppingCartId"`
	All            bool `json:"all"`
}

type ModifyInventory struct {
	Tp string `json:"tp"`
	ID uint   `json:"id"`
}
