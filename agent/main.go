package main

import (
	"monitor/agent/jvm"
)
func main() {
	go jvm.Start()
}
