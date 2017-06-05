package main

import (
"monitor/agent/jvm"
"os"
	"monitor/agent/network"
)

func main() {
	switch os.Args[1] {
	case "server":
		network.StartServer()
	case "agent":
		jvm.Start()
	}
}
