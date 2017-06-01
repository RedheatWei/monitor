package jvm

import (
	"monitor/agent/base"
	"fmt"
)
func getArgs()(*Jolokia,[]string){
	jok := &Jolokia{base.ReadConfig("jvm","jolokiapath"),base.ReadConfig("jvm","jolokianame"),base.ReadConfig("jvm","portstart")}
	pid_slice := GetPid(jok)
	return jok,pid_slice
}
//获取绑定端口
func getBindPort(jok *Jolokia,pid_slice []string) []string{
	portlist := GetPort(jok,pid_slice)
	return portlist
}
//启动
func Start(method string) {
	jok,pid_slice := getArgs()
	if method=="start"{
		StartJok(jok,pid_slice)
		portlist := getBindPort(jok,pid_slice)
		fmt.Println(portlist)
		GetRuntime(portlist)
	}else{
		StopJok(jok,pid_slice)
	}
}
