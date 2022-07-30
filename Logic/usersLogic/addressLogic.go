package usersLogic

import (
	commonModels "GO-Store/Models/common"
	"GO-Store/Models/usersModel/address"
	"GO-Store/Utils/conversion"
	"GO-Store/Utils/validator"
	"fmt"
	"github.com/sirupsen/logrus"
)

func SetAddress(userID uint, data *address.SetAddressStruct) (results interface{}, err error) {
	logrus.Info("访问了SetAddress")
	if data.ID <= 0 {
		if !validator.VerifyMobileFormat(data.Telephone) {
			return nil, fmt.Errorf("手机号格式错误")
		}
		addressInfo := address.UserAddress{
			UserId:    userID,
			Contact:   data.Contact,
			Telephone: data.Telephone,
			Province:  data.Province,
			City:      data.City,
			District:  data.District,
			Address:   data.Address,
			IsDefault: conversion.BoolTurnInt8(data.Default),
		}
		if data.Default {
			addressInfo.AllDefaultTOFalse()
		}
		if !addressInfo.Create() {
			return nil, fmt.Errorf("添加失败")
		}
		return "添加成功", nil
	} else {
		//更新地址信息
		if !validator.VerifyMobileFormat(data.Telephone) {
			return nil, fmt.Errorf("手机号格式错误")
		}

		addressInfo := address.UserAddress{
			PublicModel: commonModels.PublicModel{
				ID: data.ID,
			},
			UserId:    userID,
			Contact:   data.Contact,
			Telephone: data.Telephone,
			Province:  data.Province,
			City:      data.City,
			District:  data.District,
			Address:   data.Address,
			IsDefault: conversion.BoolTurnInt8(data.Default),
		}

		if data.Default {
			addressInfo.AllDefaultTOFalse()
		}
		if !addressInfo.Update() {
			return nil, fmt.Errorf("更新失败")
		}
		return "更新成功", nil
	}

}
func GetAddressList(userID uint) (results interface{}, err error) {
	addressList := new(address.UserAddressList)
	if !addressList.SelectById(userID) {
		return nil, fmt.Errorf("查询失败")
	}
	return addressList, nil
}
