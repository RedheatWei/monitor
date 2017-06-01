package jvm

import (
	"monitor/agent/base"
)
//启动jok
func startJok() (*Jolokia,[]string){
	jok := &Jolokia{base.ReadConfig("jvm","jolokiapath"),base.ReadConfig("jvm","jolokianame"),base.ReadConfig("jvm","portstart")}
	pid_slice := GetPid(jok)
	StartJok(jok,pid_slice)
	return jok,pid_slice
}
//获取绑定端口
func getBindPort(jok *Jolokia,pid_slice []string) []string{
	portlist := GetPort(jok,pid_slice)
	return portlist
}
//启动
func Start() {
	jok,pid_slice := startJok()
	portlist := getBindPort(jok,pid_slice)

}
