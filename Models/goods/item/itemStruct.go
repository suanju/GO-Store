package item

import (
	"GO-Store/Utils/conversion"
)

type ItemsResponse struct {
	ID           uint     `json:"id"`
	Image        string   `json:"image"`
	GoodsId      uint     `json:"goods_id"`
	SpecValueIds []string `json:"spec_value_ids"`
	SpecValueStr []string `json:"spec_value_str"`
	Price        float64  `json:"price" `
	Stock        int64    `json:"stock"`
}

type ItemsResponseList []ItemsResponse

func (l *List) Response() []ItemsResponse {
	var list []ItemsResponse

	for _, v := range *l {
		list = append(list, ItemsResponse{
			ID:           v.ID,
			Image:        conversion.FormattingSrc(v.Image),
			GoodsId:      v.GoodsId,
			SpecValueIds: conversion.StringConversionMap(v.SpecValueIds),
			SpecValueStr: conversion.StringConversionMap(v.SpecValueStr),
			Price:        v.Price,
			Stock:        v.Stock,
		})
	}
	return list

}
