package jvm

import (
	"monitor/agent/base"
	"fmt"
	"os"
)
var PortBinding []string
func getArgs()(*Jolokia,[]string){
	jok := &Jolokia{base.ReadConfig("jvm","jolokiapath"),base.ReadConfig("jvm","jolokianame"),base.ReadConfig("jvm","portstart")}
	pid_slice := GetPid(jok)
	return jok,pid_slice
}
//获取绑定端口
//func getBindPort(jok *Jolokia,pid_slice []string) []string{
//	portlist := GetPort(jok,pid_slice)
//	return portlist
//}
//检查是否存在成功绑定端口
func getBaseUrl() []string{
	if len(PortBinding) !=0{
		return PortBinding
	}else{
		fmt.Println("Cannot found bind port!")
		os.Exit(1)
		return []string{}
	}
}
//启动
func Start(method string) {
	jok,pid_slice := getArgs()
	if len(pid_slice) == 0{
		fmt.Println("Cannot found java process!")
		os.Exit(1)
	}
	switch method {
	case "stop":
		StopJok(jok,pid_slice)
	case "get":
		StartJok(jok,pid_slice)
		baseUrl := getBaseUrl()
		ch := make(chan []string)
		go AccceptGet(baseUrl)
		ch <- getBaseUrl()
	}
}
