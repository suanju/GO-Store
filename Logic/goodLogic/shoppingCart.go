package goodLogic

import (
	shoppingCart2 "GO-Store/Models/usersModel/shoppingCart"
	"fmt"
)

func AddShoppingCart(data *shoppingCart2.AddShoppingCart, uid uint) (results interface{}, err error) {
	//添加购物车
	shoppingCart := shoppingCart2.ShoppingCart{
		UserID:  uid,
		GoodsId: data.GoodsID,
		ItemID:  data.ItemId,
		Number:  data.Number,
	}
	if !shoppingCart.AddGood() {
		return nil, fmt.Errorf("加入购物车失败")
	}
	return "加入成功", nil
}

func GetShoppingCart(data *shoppingCart2.GetShoppingCart, uid uint) (results interface{}, err error) {
	//查询当前用用户的购物车
	shoppingCartList := new(shoppingCart2.List)
	if err := shoppingCartList.GetList(uid); err != nil {
		return fmt.Errorf("查询失败"), nil
	}
	results = shoppingCartList.Response()
	return results, nil
}
