package handler

import (
	"encoding/json"
	"fmt"
)
type jsoninfo struct {
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
func ToJson(rec []byte) jsoninfo{
	var js jsoninfo
	str := string(rec)
	json.Unmarshal([]byte(str),&js)
	fmt.Println(js)
	return js
}