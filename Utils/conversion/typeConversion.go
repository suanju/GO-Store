package conversion

import "strings"

func StringConversionMap(s string) []string {
	list := strings.Split(s, ",")
	return list
}

func StringImgConversionMap(s string) []string {
	list := strings.Split(s, ",")
	for k, v := range list {
		list[k] = FormattingSrc(v)
	}
	return list
}

//BoolTurnInt8 布尔类型转int8
func BoolTurnInt8(is bool) int8 {
	if is {
		return 1
	} else {
		return 0
	}
}
