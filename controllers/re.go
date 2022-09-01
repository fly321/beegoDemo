package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ReController struct {
	beego.Controller
}

func (c *ReController) RgTest() {
	cid := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("id is " + cid)
}
