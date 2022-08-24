package goodLogic

import (
	"GO-Store/Models/goodsModel/tabs"
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
