package jvm

import (
	"monitor/base"
	"fmt"
	"os"
	"strconv"
)
var PortBinding []string
var Frequency int64; var(
	//met string = "read"
	args = []string{"Memory","Runtime","Threading","ClassLoading"} //OperatingSystem
)
func getArgs()(*Jolokia,[]string){
	jok := &Jolokia{base.ReadAgentConfig("jvm","jolokiapath"),base.ReadAgentConfig("jvm","jolokianame"),base.ReadAgentConfig("jvm","portstart")}
	pid_slice := GetPid(jok)
	return jok,pid_slice
}
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
func Start() {

	jok,pid_slice := getArgs()
	Frequency,_ = strconv.ParseInt(base.ReadAgentConfig("default","frequency"),10,64)
	if len(pid_slice) == 0{
		fmt.Println("Cannot found java process!")
		os.Exit(1)
	}

	StartJok(jok,pid_slice)
	baseUrl := getBaseUrl()
	ch := make(chan []string)
	go CollectJvmInfo(baseUrl,args)
	ch <- baseUrl
}
