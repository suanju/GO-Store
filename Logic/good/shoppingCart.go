package good

import (
	shoppingCart2 "GO-Store/Models/users/shoppingCart"
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

func GetShoppingCart(uid uint) (results interface{}, err error) {
	//查询当前用用户的购物车
	shoppingCartList := new(shoppingCart2.List)
	if err := shoppingCartList.GetList(uid); err != nil {
		return fmt.Errorf("查询失败"), nil
	}
	results = shoppingCartList.Response()
	return results, nil
}

func DelShoppingCart(data *shoppingCart2.DelShoppingCart, uid uint) (results interface{}, err error) {
	//判断是否为清空购物车
	shoppingCart := new(shoppingCart2.ShoppingCart)
	if data.All {
		if err := shoppingCart.Empty(uid); err != nil {
			return fmt.Errorf("清空失败"), nil
		}
		return "清空成功", nil
	} else {
		//删除单条数据
		if err := shoppingCart.DelByID(data.ShoppingCartID); err != nil {
			return fmt.Errorf("删除失败"), nil
		}
		return "删除成功", nil
	}
}

func ModifyInventory(data *shoppingCart2.ModifyInventory, uid uint) (results interface{}, err error) {
	//判断是否为清空购物车
	//判断传参是否正确
	//tpArr := []string{"+", "-"}
	shoppingCartList := new(shoppingCart2.ShoppingCart)
	if err := shoppingCartList.ModifyInventory(uid, data.ID, data.Tp); err != nil {
		return err, nil
	}
	return "修改成功", nil
}
