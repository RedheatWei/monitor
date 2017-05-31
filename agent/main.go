package main

import (
	"monitor/agent/jvm"
)

func main() {
	jok := &jvm.Jolokia{"/usr/local/jolokia/","jolokia-jvm-1.3.6-agent.jar"}
	pid_slice := jvm.GetPid(jok)
	jvm.StartJok(jok,pid_slice)
}
