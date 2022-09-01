package goods

import (
	"GO-Store/Models/common"
	"GO-Store/Models/goods/comments"
	"GO-Store/Models/goods/item"
	"GO-Store/Models/goods/spec"
	spec2 "GO-Store/Models/goods/specValue"
	conversion "GO-Store/Utils/conversion"
)

// GetClassifyByLevelStruct 获取分类列表请求结构体
type GetClassifyByLevelStruct struct {
	common.PageInfo
	Level int64 `json:"level" binding:"required"`
	Pid   int64 `json:"pid" binding:"required"`
}

// GetGoodInfoStruct 获取商品详情
type GetGoodInfoStruct struct {
	Id int64 `json:"id"`
}

type GetCollectionStruct struct {
	Uid          uint `json:"uid"`
	GoodsID      uint `json:"goods_id"`
	IsCollection bool `json:"is_collection"`
}

/*上面接口结构体---------------------------------------------------------------下面返回数据结构体*/

type GoodListResponse struct {
	ID    uint    `json:"id"`
	Image string  `json:"image"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Sales int64   `json:"sales"`
}

type GetTabsInfo struct {
	List GoodList `json:"list"`
	Size int      `json:"size"`
}

type GoodResponse struct {
	ID           uint                   `json:"id"`
	Name         string                 `json:"name"`
	Image        string                 `json:"image"`
	ShufflingImg []string               `json:"shuffling_img"`
	Video        string                 `json:"video"`
	Remark       string                 `json:"remark"`
	SalesActual  int64                  `json:"sales_actual"`
	Clicks       int64                  `json:"clicks"`
	Price        float64                `json:"price"`
	Stock        int64                  `json:"stock"`
	Content      string                 `json:"content"`
	Spec         spec.List              `json:"spec"`
	SpecValue    spec2.ValueList        `json:"spec_value"`
	Item         item.ItemsResponseList `json:"item"`
	SeverList    []string               `json:"sever_list"`
	Comments     comments.ListResponse  `json:"comments"`
	Like         bool                   `json:"like" `
}

func (g *Goods) Response() GoodResponse {
	info := GoodResponse{
		ID:           g.ID,
		Name:         g.Name,
		Image:        conversion.FormattingSrc(g.Image),
		ShufflingImg: conversion.StringImgConversionMap(g.ShufflingImg),
		Video:        g.Video,
		Remark:       g.Remark,
		SalesActual:  g.SalesActual,
		Clicks:       g.Clicks,
		Price:        g.Price,
		Stock:        g.Stock,
		Spec:         g.Spec,
		Content:      g.Content,
		SpecValue:    g.SpecValue.Response(),
		Item:         g.Item.Response(),
		SeverList:    g.SeverList.ConversionRsp(),
		Comments:     g.Comments.Response(),
		Like:         g.Like.IsLike(),
	}
	return info
}
func (l *GoodList) Response() []GoodListResponse {
	var list []GoodListResponse
	for _, v := range *l {
		list = append(list, GoodListResponse{
			ID:    v.ID,
			Image: conversion.FormattingSrc(v.Image),
			Title: v.Name,
			Price: v.Price,
			Sales: v.SalesVirtual,
		})
	}
	return list

}
