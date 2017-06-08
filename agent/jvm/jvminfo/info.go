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

func GetInfo(baseUrl string) Info{
	_,res := base.HttpGet(baseUrl)
	var info Info
	json.Unmarshal(res,&info)
	return info
}
