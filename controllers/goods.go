package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

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
	//id := c.GetString("id")
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id is not int")
		return
	}
	c.Ctx.WriteString(strconv.Itoa(id))
}

func (c *GoodsController) DoAdd() {
	name := c.GetString("name", "defaultName")
	c.Ctx.WriteString(name)
}
