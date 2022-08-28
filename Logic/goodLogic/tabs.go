package goodLogic

import (
	"GO-Store/Models/goodsModel/goods"
	"GO-Store/Models/goodsModel/tabs"
	"GO-Store/Utils/calculate"
)

func GetTabs(data *tabs.GetTabsStruct) (results interface{}, err error) {
	//获取tabs
	tabsList := new(tabs.TabList)
	err = tabsList.GetAllList()
	if err != nil {
		return err, nil
	}
	return tabsList, nil
}
func GetTabsInfo(data *tabs.GetTabsInfoStruct) (results interface{}, err error) {
	//获取tabs信息
	id := data.TabsId
	dataInfo := new(goods.GetTabsInfo)
	tabsInfo := new(tabs.Tabs)
	goodList := new(goods.GoodList)
	tabsInfo.ID = id
	err = tabsInfo.GetInfoById()
	if err != nil {
		return nil, err
	}
	//取值
	err = goodList.SelectByTab(dataInfo, tabsInfo.ID, data.PageNum, data.PageSize)
	if err != nil {
		return nil, err
	}
	res := new(tabs.GetTabsInfoResponseStruct)
	res.Size = dataInfo.Size
	res.List = dataInfo.List.Response()
	//根据总数判断页数
	res.LastPage = calculate.ComputationalPages(dataInfo.Size, data.PageSize)

	return res, nil
}
