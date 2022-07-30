package goodLogic

import (
	"GO-Store/Models/goodsModel/goods"
)

func GetClassifyByLevel(data *goods.GetClassifyByLevelStruct) (results interface{}, err error) {
	data.Init()
	goodsList := new(goods.GoodList)
	if err := goodsList.SelectByLevel(data.Level, data.Pid, data.PageInfo); err != nil {
		return nil, err
	}
	results = goodsList.Response()
	return results, nil
}
func GetGoodInfo(data *goods.GetGoodInfoStruct) (results interface{}, err error) {
	goodInfo := new(goods.Goods)
	if err := goodInfo.SelectById(data.Id); err != nil {
		return err, err
	}
	//点击量
	goodInfo.ClicksAdd()
	results = goodInfo.Response()
	return results, nil
}
