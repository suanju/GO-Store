package address

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/common"
)

//UserAddress 表结构体
type UserAddress struct {
	common.PublicModel
	UserId    uint   `json:"user_id" gorm:"user_id"`
	Contact   string `json:"contact" gorm:"contact"`
	Telephone string `json:"telephone" gorm:"telephone"`
	Province  string `json:"province" gorm:"province"`
	City      string `json:"city" gorm:"city"`
	District  string `json:"district" gorm:"district"`
	Address   string `json:"address" gorm:"address"`
	PostCode  int    `json:"-" gorm:"post_code"`
	Longitude string `json:"-" gorm:"longitude"`
	Latitude  string `json:"-" gorm:"latitude"`
	IsDefault int8   `json:"is_default" gorm:"is_default"`
}

type UserAddressList []UserAddress

func (UserAddress) TableName() string {
	return "ml_user_address"
}

//Update 更新数据
func (us *UserAddress) Update() bool {
	err := Mysql.Db.Model(&us).Updates(&us).Error
	if err != nil {
		return false
	}
	return true
}

//Create 添加数据
func (us *UserAddress) Create() bool {
	err := Mysql.Db.Create(&us).Error
	if err != nil {
		return false
	}
	return true
}

//SelectById 获取用户地址列表
func (us *UserAddressList) SelectById(id uint) bool {
	err := Mysql.Db.Where("user_id", id).Find(&us).Error
	if err != nil {
		return false
	}
	return true
}

func (us *UserAddress) AllDefaultTOFalse() bool {
	userAdd := new(UserAddress)
	err := Mysql.Db.Model(&userAdd).Where("user_id = ?", us.UserId).Update("is_default", 0).Error
	if err != nil {
		return false
	}
	return true
}
