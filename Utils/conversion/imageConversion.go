package conversion

import "fmt"

//FormattingSrc 图片处理相关
func FormattingSrc(src string) string {
	return fmt.Sprintf("http://127.0.0.1:8080%s", src)
}
