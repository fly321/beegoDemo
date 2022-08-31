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
}
