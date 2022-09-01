package shoppingCart

import "GO-Store/Utils/conversion"

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

//返回需要

//LResponse 返回结构体
type LResponse struct {
	Id      uint    `json:"id"`
	Image   string  `json:"image"`
	AttrVal string  `json:"attr_val"`
	Stock   int64   `json:"stock"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	Number  int64   `json:"number"`
}
type ListResponse []LResponse

func (ls List) Response() ListResponse {
	var list ListResponse
	for _, v := range ls {
		list = append(list, LResponse{
			Id:      v.ID,
			Image:   conversion.FormattingSrc(v.Item.Image),
			AttrVal: v.Item.SpecValueStr,
			Stock:   v.Item.Stock,
			Title:   v.GoodsInfo.Name,
			Price:   v.Item.Price,
			Number:  v.Number,
		})
	}
	return list

}
