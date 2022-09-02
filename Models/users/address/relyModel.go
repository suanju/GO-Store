package address

import "GO-Store/Databases/Mysql"

//Region 表结构体
type Region struct {
	ID       int64  `json:"id" gorm:"id"`
	ParentId int64  `json:"parent_id" gorm:"parent_id"`
	Level    int64  `json:"level" gorm:"level"`
	Name     string `json:"name" gorm:"name"`
	CityCode string `json:"city_code" gorm:"city_code"`
	CityId   int64  `json:"city_id"`
}

func (Region) TableName() string {
	return "ml_dev_region"
}

type RegionList []Region

type OutputRegionList []OutputRegion

//SelectByParentId 获取用户地址列表
func (rl *RegionList) SelectByParentId(id int64) (list OutputRegionList, err error) {
	err = Mysql.Db.Where("parent_id", id).Find(&rl).Error
	if err != nil {
		return nil, err
	}
	list = rl.FormattedOutput()
	return list, nil
}
