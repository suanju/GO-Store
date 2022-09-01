package banner

import (
	"GO-Store/Models/common"
	LoginCommon "GO-Store/Utils/conversion"
)

// BannersHome 主页轮播图表结构
type BannersHome struct {
	common.PublicModel
	Src        string `json:"src" gorm:"src"`
	SrcType    int64  `json:"src_type" gorm:"src_type"`
	Background string `json:"background" gorm:"background"`
	Type       int64  `json:"type" gorm:"type"`
	JumpId     int64  `json:"jump_id" gorm:"jump_id"`
}

func (BannersHome) TableName() string {
	return "ml_banner_home"
}

// BannersHomeList 主页轮播图表结构数组
type BannersHomeList []BannersHome

//FormattingSrc 主页轮播图表结构数组格式化Src
func (l BannersHomeList) FormattingSrc() {
	for k, v := range l {
		l[k].Src = LoginCommon.FormattingSrc(v.Src)
	}
}
