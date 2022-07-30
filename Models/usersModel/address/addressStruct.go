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

/*上面接口结构体---------------------------------------------------------------下面返回数据结构体*/
