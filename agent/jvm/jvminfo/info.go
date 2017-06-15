package jvminfo

import (
	"encoding/json"
	"monitor/agent/base"
)

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
//获取jvm基本信息
func GetInfo(baseUrl string) (n int,info Info){
	n,res := base.HttpGet(baseUrl)
	if n==0{
		var info Info
		json.Unmarshal(res,&info)
		return 0,info
	}else {
		var info Info
		return 1,info
	}

}
