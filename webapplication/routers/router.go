package routers

import (
	"monitor/webapplication/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/jvm/GET", &controllers.JvmgetController{})
}
