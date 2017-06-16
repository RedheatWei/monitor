package main

import (
	//"monitor/agent/jvm"
	"monitor/agent/system"
)
//主函数
func main() {
	ch := make(chan string)
	//go jvm.Start()
	go system.Start()
	ch <- "a"
}