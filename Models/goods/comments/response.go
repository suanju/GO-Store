package comments

import (
	"GO-Store/Models/common"
	"GO-Store/Utils/conversion"
)

type CommentUserInfoResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Photo    string `json:"photo"`
}

type Response struct {
	common.PublicModel
	GoodsId     uint     `json:"goods_id" `
	UserID      uint     `json:"user_id" `
	SpecValueId []string `json:"spec_value_id" `
	Text        string   `json:"text" `
	Image       []string `json:"image"`

	UserInfo CommentUserInfoResponse `json:"user_info"`
}

type ListResponse struct {
	Number int        `json:"number"`
	List   []Response `json:"list"`
}

func (l List) Response() ListResponse {
	var response ListResponse
	var list []Response
	for _, v := range l {
		list = append(list, Response{
			PublicModel: common.PublicModel{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			GoodsId:     v.GoodsId,
			UserID:      v.UserID,
			SpecValueId: conversion.StringConversionMap(v.SpecValueId),
			Text:        v.Text,
			Image:       conversion.StringImgConversionMap(v.Image),
			UserInfo: CommentUserInfoResponse{
				ID:       v.UserInfo.ID,
				UserName: v.UserInfo.Username,
				Photo:    conversion.FormattingSrc(v.UserInfo.Photo),
			},
		})
	}
	if len(list) > 0 {
		response.List = list[len(list)-1:]
	} else {
		response.List = list
	}

	response.Number = len(list)

	return response

}
