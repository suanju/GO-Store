package category

import "GO-Store/Utils/conversion"

type resStruct struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Pid     int64  `json:"pid"`
	Picture string `json:"picture"`
}

func (l GoodsCategoryList) Response() []resStruct {
	var list []resStruct
	for _, v := range l {
		var picture string
		if len(v.Picture) != 0 {
			picture = conversion.FormattingSrc(v.Picture)
		} else {
			picture = v.Picture
		}
		list = append(list, resStruct{
			Id:      v.ID,
			Name:    v.Name,
			Pid:     v.Pid,
			Picture: picture,
		})
	}
	return list

}
