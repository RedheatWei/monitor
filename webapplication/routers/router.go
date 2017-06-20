package routers

import (
	"monitor/webapplication/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/jvm", &controllers.JvmgetController{})
}
