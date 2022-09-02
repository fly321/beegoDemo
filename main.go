package main

import (
	_ "beegoDemo1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//beego.BConfig.WebConfig.TemplateLeft = "<<<"
	//beego.BConfig.WebConfig.TemplateRight = ">>>"
	beego.Run()
}
