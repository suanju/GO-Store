package conversion

import (
	"GO-Store/Databases"
	"fmt"
)

//FormattingSrc 图片处理相关
func FormattingSrc(src string) string {
	api := Databases.ProjectUrl
	return fmt.Sprintf("%s%s", api, src)
}
