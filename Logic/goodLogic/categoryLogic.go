package goodLogic

import (
	"GO-Store/Models/goods/category"
	"fmt"
)

func GetGoodsCategory(data *category.GetGoodsCategoryStruct) (results interface{}, err error) {
	var goodsCategoryList category.GoodsCategoryList
	if !goodsCategoryList.SelectAll() {
		return nil, fmt.Errorf("查询失败")
	}
	results = goodsCategoryList.Response()
	return results, nil
}
