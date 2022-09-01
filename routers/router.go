package routers

import (
	"beegoDemo1/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/add", &controllers.GoodsController{}, "get:Add")
	beego.Router("/goods/edit", &controllers.GoodsController{}, "get:Edit")
	beego.Router("/goods/doAdd", &controllers.GoodsController{}, "post:DoAdd")
	beego.Router("/goods/getGood", &controllers.GoodsController{}, "get:GetGood")
	beego.Router("/goods/postXml", &controllers.GoodsController{}, "post:PostXml")
	beego.Router("/goods/apiDemo/:id", &controllers.GoodsController{}, "get:ApiDemo")
	beego.Router("/rg/rgtest/:id([0-9]+).html", &controllers.ReController{}, "get:RgTest")
}
