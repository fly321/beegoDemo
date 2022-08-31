package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
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

type User struct {
	Id    int      `form:"id"`
	Name  string   `form:"name"`
	Hobby []string `form:"hobby"`
}

func (c *GoodsController) DoAdd() {
	/*name := c.GetString("name", "defaultName")
	c.Ctx.WriteString(name)*/

	user := User{}
	if err := c.ParseForm(&user); err != nil {
		c.Ctx.WriteString("parse form error")
		return
	}
	fmt.Printf("%#v\r\n", user)
	c.Ctx.WriteString(user.Name)
}

func (c *GoodsController) GetGood() {
	user := User{
		Id:    1,
		Name:  "fly",
		Hobby: []string{"sing", "dance"},
	}
	c.Data["json"] = user
	c.ServeJSON()
}

type Product struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) PostXml() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	x := string(c.Ctx.Input.RequestBody)
	log.Info(x)
	p := Product{}
	if err1 := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); err1 != nil {
		//log.Error(err1)
		c.Data["json"] = err1
		c.ServeJSON()
	} else {
		fmt.Printf("%#v", p)
		c.Data["json"] = p
		c.ServeJSON()
	}

}
