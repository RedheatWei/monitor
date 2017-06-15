package jvminfo

import "encoding/json"

//"Threading","ClassLoading" OperatingSystem
type OperatingSystem struct {
	//Request runtimeRequest `json:"request"`
	Value operatingSystemValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type operatingSystemValue struct {
	OpenFileDescriptorCount int32 `json:"OpenFileDescriptorCount"`
	CommittedVirtualMemorySize int64 `json:"CommittedVirtualMemorySize"`
	FreePhysicalMemorySize int64 `json:"FreePhysicalMemorySize"`
	SystemLoadAverage float64 `json:"SystemLoadAverage"`
	ProcessCpuLoad float64 `json:"ProcessCpuLoad"`
	FreeSwapSpaceSize int64 `json:"FreeSwapSpaceSize"`
	TotalPhysicalMemorySize int64 `json:"TotalPhysicalMemorySize"`
	TotalSwapSpaceSize int64 `json:"TotalSwapSpaceSize"`
	AvailableProcessors int32 `json:"AvailableProcessors"`
}
//获取系统信息,此方法并未使用
func GetOperatingSystem(baseUrl string) OperatingSystem{
	res := ResGet(baseUrl,"OperatingSystem")
	var operatingSystem OperatingSystem
	json.Unmarshal(res,&operatingSystem)
	return operatingSystem
}

