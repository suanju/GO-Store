package usersLogic

import (
	"GO-Store/Models/users/address"
	"fmt"
)

func GetAddressTable(data *address.RegionReceiveStruct) (results interface{}, err error) {
	region := new(address.RegionList)
	results, err = region.SelectByParentId(data.CityId)
	if err != nil {
		return nil, fmt.Errorf("获取失败")
	}
	return results, nil
}
