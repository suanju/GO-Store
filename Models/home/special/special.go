package special

import (
	"GO-Store/Models/common"
	LoginCommon "GO-Store/Utils/conversion"
)

// SpecialsHome 主页会场表结构
type SpecialsHome struct {
	common.PublicModel
	Src  string `json:"src" gorm:"src"`
	Text string `json:"text" gorm:"text"`
}

func (SpecialsHome) TableName() string {
	return "ml_special"
}

// SpecialsHomeList 主页会场表结构数组
type SpecialsHomeList []SpecialsHome

//FormattingSrc 主页会场表结构数组格式化Src
func (l SpecialsHomeList) FormattingSrc() {
	for k, v := range l {
		l[k].Src = LoginCommon.FormattingSrc(v.Src)
	}
}
