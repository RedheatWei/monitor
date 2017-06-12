package controllers

import (
	"github.com/astaxie/beego"
	"monitor/webapplication/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}
func (c *MainController) Get() {
	c.Data["Website"] = "golang.org"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *TestController) Get(){
	c.Data["test"] = "this is a test"
	r := models.Sel()
	fmt.Println(r)
	c.TplName = "test.html"
}