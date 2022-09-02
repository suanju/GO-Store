package home

import (
	"GO-Store/Databases/Mysql"
	"GO-Store/Models/home/banner"
	"GO-Store/Models/home/special"
)

func GetBannerList(data *banner.GetBannerListStruct) (results interface{}, err error) {
	var bannerList banner.BannersHomeList
	err = Mysql.Db.Find(&bannerList).Error
	if err != nil {
		return nil, err
	}
	bannerList.FormattingSrc()
	return bannerList, nil
}

func GetSpecialList(data *special.GetSpecialListStruct) (results interface{}, err error) {
	var specialList special.SpecialsHomeList
	err = Mysql.Db.Find(&specialList).Error
	if err != nil {
		return nil, err
	}
	specialList.FormattingSrc()
	return specialList, nil
}
