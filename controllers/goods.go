package controllers

import beego "github.com/beego/beego/v2/server/web"

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "fly321"
	c.TplName = "goods.html"
}

func (c *GoodsController) Add() {
	c.Ctx.WriteString("add goods")
}

func (c *GoodsController) Edit() {
	id := c.GetString("id")
	c.Ctx.WriteString("id:" + id)
}

func (c *GoodsController) DoAdd() {
	name := c.GetString("name", "defaultName")
	c.Ctx.WriteString(name)
}
