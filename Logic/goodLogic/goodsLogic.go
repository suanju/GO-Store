package goodLogic

import (
	"GO-Store/Models/goods/goods"
	"GO-Store/Models/users/likeGoods"
	"fmt"
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

func GetGoodInfo(data *goods.GetGoodInfoStruct, uid uint) (results interface{}, err error) {
	goodInfo := new(goods.Goods)
	if err := goodInfo.SelectById(data.Id, uid); err != nil {
		return err, err
	}
	//点击量
	goodInfo.ClicksAdd()
	results = goodInfo.Response()
	return results, nil
}

func Collection(data *goods.GetCollectionStruct, uid uint) (results interface{}, err error) {
	//收藏商品
	like := new(likeGoods.LikeGoods)
	if !like.Like(uid, data.GoodsID) {
		return nil, fmt.Errorf("收藏失败")
	}
	return "收藏成功", nil
}
