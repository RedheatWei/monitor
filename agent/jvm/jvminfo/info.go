package jvminfo

import (
	"encoding/json"
	"monitor/base"
)
//"Threading","ClassLoading" OperatingSystem
type Info struct {
	//Request runtimeRequest `json:"request"`
	Value classLoadingValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type InfoValue struct {
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
}

func GetInfo(baseUrl string) Info{
	_,res := base.HttpGet(baseUrl)
	var info Info
	json.Unmarshal(res,&info)
	return info
}
