package goods

import "GO-Store/Models/common"

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
