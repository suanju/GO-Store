package main

import (
	_ "GO-Store/Databases/Mysql"
	_ "GO-Store/Databases/Redis"
	"GO-Store/Router"
)

func main() {
	Router.InitRouter()
}
