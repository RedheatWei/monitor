package main

import (
	"monitor/agent/jvm"
)
func main() {
	ch := make(chan string)
	go jvm.Start()
	ch <- "a"
}
