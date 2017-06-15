package jvminfo
import "encoding/json"

type Threading struct {
	//Request runtimeRequest `json:"request"`
	Value threadingValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type threadingValue struct {
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime int64 `json:"CurrentThreadCpuTime"`
	ThreadCount int32 `json:"ThreadCount"`
	DaemonThreadCount int32 `json:"DaemonThreadCount"`
}
//获取jvm线程信息
func GetThreading(baseUrl string) Threading{
	res := ResGet(baseUrl,"Threading")
	var threading Threading
	json.Unmarshal(res,&threading)
	return threading
}
