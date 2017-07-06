package main

import (
	"monitor/agent/jvm"
	"monitor/agent/system"
	"monitor/agent/base"
)
//主函数
func main() {
	conf := base.GetConfig()
	ch := make(chan string)
	if conf.Jvm.On=="true"{
		go jvm.Start()
	}
	go system.Start()
	ch <- "a"
}