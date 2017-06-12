package main

import (
	_ "monitor/webapplication/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

