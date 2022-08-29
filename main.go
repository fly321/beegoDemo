package main

import (
	_ "beegoDemo1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
