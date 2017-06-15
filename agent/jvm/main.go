package jvm

import (
	"monitor/agent/base"
	"fmt"
	"os"
	"strconv"
)
var PortBinding []string
var AgentConfig base.AgentConfig
var Frequency int64
//读取配置文件
func init()  {
	AgentConfig = base.GetConfig()
}
func getArgs()(*Jolokia,[]string){
	jok := &Jolokia{AgentConfig.Jvm.Jolokiapath,AgentConfig.Jvm.Jolokianame,AgentConfig.Jvm.Portstart}
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
	Frequency,_ = strconv.ParseInt(AgentConfig.Default.Frequency,10,64)
	if len(pid_slice) == 0{
		fmt.Println("Cannot found java process!")
		os.Exit(1)
	}

	StartJok(jok,pid_slice)
	baseUrl := getBaseUrl()
	CollectJvmInfo(baseUrl)
}
