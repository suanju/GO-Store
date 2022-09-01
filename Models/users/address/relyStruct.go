package address

//RegionReceiveStruct 地区表请求结构提
type RegionReceiveStruct struct {
	CityId int64 `json:"city_id" binding:"required"`
}
type OutputRegion struct {
	CityId int64  `json:"city_id"`
	Name   string `json:"name"`
}
