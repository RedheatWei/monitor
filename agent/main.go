package main

import (
	//"monitor/agent/jvm"
	"monitor/agent/config"
	"fmt"
	"monitor/agent/conf"
)

func main() {
	var Config = new(conf.Config)
	Config.InitConfig("config/agent.conf")
	fmt.Println(Config.Read("jvm", "jolokiapath"))
	fmt.Println(Config.Read("jvm", "portstart"))
	//jok := &jvm.Jolokia{"/usr/local/jolokia/","jolokia-jvm-1.3.6-agent.jar"}
	//pid_slice := jvm.GetPid(jok)
	//jvm.StartJok(jok,pid_slice)
}
