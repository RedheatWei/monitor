package jvminfo

import (
	"encoding/json"
	"monitor/base"
)
//"Threading","ClassLoading" OperatingSystem
type Info struct {
	//Request runtimeRequest `json:"request"`
	Value infoValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type infoValue struct {
	Config config `json:"config"`
}
type config struct {
	AgentId string `json:"agentId"`
}

func GetInfo(baseUrl string) (n int,info Info){
	n,res := base.HttpGet(baseUrl)
	var info Info
	if n==0{
		json.Unmarshal(res,&info)
		return 0,info
	}else {
		return 1,info
	}

}
