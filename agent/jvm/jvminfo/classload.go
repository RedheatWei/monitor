package jvminfo

import "encoding/json"

//"Threading","ClassLoading" OperatingSystem
type ClassLoading struct {
	//Request runtimeRequest `json:"request"`
	Value classLoadingValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type classLoadingValue struct {
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
}

func GetClassLoading(baseUrl string) ClassLoading{
	res := ResGet(baseUrl,"ClassLoading")
	var classLoading ClassLoading
	json.Unmarshal(res,&classLoading)
	return classLoading
}
