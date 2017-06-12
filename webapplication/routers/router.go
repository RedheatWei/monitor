package routers

import (
	"monitor/webapplication/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/text", &controllers.MainController{})
}
