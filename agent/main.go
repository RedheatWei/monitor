package main

import (
	//"monitor/agent/jvm"
	"monitor/agent/config"
	"fmt"
)

func main() {
	Config := new(conf.Config)
	Config.InitConfig("config/agent.conf")
	fmt.Println(Config.Read("jvm", "jolokiapath"))
	fmt.Println(Config.Read("jvm", "portstart"))
	fmt.Printf("%v", Config.Mymap)
	//jok := &jvm.Jolokia{"/usr/local/jolokia/","jolokia-jvm-1.3.6-agent.jar"}
	//pid_slice := jvm.GetPid(jok)
	//jvm.StartJok(jok,pid_slice)
}
