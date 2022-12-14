package address

//SetAddressStruct 地区表请求结 构提
type SetAddressStruct struct {
	ID        uint   `json:"id"`
	Contact   string `json:"contact" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Province  string `json:"province" binding:"required"`
	City      string `json:"city" binding:"required"`
	District  string `json:"district" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Default   bool   `json:"is_default"`
}
type GetAddressListStruct struct {
}

//RegionReceiveStruct 地区表请求结构提
type RegionReceiveStruct struct {
	CityId int64 `json:"city_id" binding:"required"`
}
type OutputRegion struct {
	CityId int64  `json:"city_id"`
	Name   string `json:"name"`
}
