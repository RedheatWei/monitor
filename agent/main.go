package main

import (
	"monitor/agent/jvm"
	"os"
)

func main() {
	jvm.Start(os.Args[1])
}
