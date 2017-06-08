package base

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"monitor/conf"
)

type JsonInfo struct {
	//class
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit int64 `json:"init"`
	HeapMemoryUsageCommitted int64 `json:"committed"`
	HeapMemoryUsageMax int64 `json:"max"`
	HeapMemoryUsageUsed int64 `json:"used"`
	NonHeapMemoryUsageInit int64 `json:"init"`
	NonHeapMemoryUsageCommitted int64 `json:"committed"`
	NonHeapMemoryUsageMax int64 `json:"max"`
	NonHeapMemoryUsageUsed int64 `json:"used"`
	//rumtime
	ClassPath string `json:"ClassPath"`
	Uptime int64 `json:"Uptime"`
	//thread
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime int64 `json:"CurrentThreadCpuTime"`
	ThreadCount int32 `json:"ThreadCount"`
	DaemonThreadCount int32 `json:"DaemonThreadCount"`
	//base
	AgentId string `json:"agentId"`
	TimeStamp int64 `json:"TimeStamp"`
}

//get方法
func HttpGet(url string) (n int,res []byte){
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1,[]byte{}
	}else{
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		return 0,body
	}
}
//读取配置文件
func ReadAgentConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("../conf/agent.conf")
	return config.Read(mod,par)
}
//读取服务端配置文件
func ReadServerConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("../conf/server.conf")
	return config.Read(mod,par)
}