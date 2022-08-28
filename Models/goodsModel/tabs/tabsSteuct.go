package tabs

// GetTabsStruct 获取tabs
type GetTabsStruct struct {
}

// GetTabsInfoStruct 获取tabs当中的商品
type GetTabsInfoStruct struct {
	PageNum  int  `json:"pageNum"`
	PageSize int  `json:"pageSize"`
	TabsId   uint `json:"tabId"`
}

type GetTabsInfoResponseStruct struct {
	LastPage int         `json:"lastPage"`
	List     interface{} `json:"list"`
	Size     int         `json:"size"`
}
