package main

import (
	"monitor/agent/conf"
	"monitor/agent/jvm"
	"fmt"
)

func main() {
	jok := &jvm.Jolokia{ReadConfig("jvm","jolokiapath"),ReadConfig("jvm","jolokianame"),ReadConfig("jvm","portstart")}
	pid_slice := jvm.GetPid(jok)
	//jvm.StartJok(jok,pid_slice)
	//port := "18005"
	//url := "http://127.0.0.1:"+port+"/jolokia/"
	portlist := jvm.GetPort(jok,pid_slice)
	fmt.Println(portlist)
}
func ReadConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/agent.conf")
	return config.Read(mod,par)
}