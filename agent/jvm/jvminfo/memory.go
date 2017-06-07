package jvminfo

import (
	"encoding/json"
	"monitor/agent/jvm"
)
//"Memory"
type Memory struct {
	Request memoryRequest `json:"request"`
	Value memoryValue `json:"value"`
	TimeStamp int32 `json:"timestamp"`
	Status int32 `json:"status"`
}
type memoryusage struct {
	Init int64 `json:"init"`
	Committed int64 `json:"committed"`
	Max int64 `json:"max"`
	Used int64 `json:"used"`
}
type memoryValue struct {
	ObjectPendingFinalizationCount int `json:"ObjectPendingFinalizationCount"`
	Verbose bool `json:"Verbose"`
	HeapMemoryUsage memoryusage `json:"HeapMemoryUsage"`
	NonHeapMemoryUsage memoryusage `json:"NonHeapMemoryUsage"`
	ObjectName objectname `json:"ObjectName"`
}
type memoryRequest struct {
	Mbean string `json:"mbean"`
	Type string `json:"type"`
}
type objectname struct {
	ObjectName string `json:"objectName"`
}
func GetMemory(baseUrl string) Memory{
	res := jvm.ResGet(baseUrl,"Memory")
	var memoryUsage Memory
	json.Unmarshal(res,&memoryUsage)
	return memoryUsage
}