package jvminfo

import "encoding/json"

//"Runtime","Threading","ClassLoading" OperatingSystem
type Runtime struct {
	//Request runtimeRequest `json:"request"`
	Value runtimeValue `json:"value"`
	TimeStamp int64 `json:"timestamp"`
	Status int32 `json:"status"`
}
type runtimeRequest struct {
	Mbean string `json:"mbean"`
	Type string `json:"type"`
}
type runtimeValue struct {
	//SpecVendor string `json:"SpecVendor"`
	ClassPath string `json:"ClassPath"`
	//InputArguments []string `json:"InputArguments"`
	Uptime int64 `json:"Uptime"`
}

func GetRuntime(baseUrl string) Runtime{
	res := ResGet(baseUrl,"Runtime")
	var runtime Runtime
	json.Unmarshal(res,&runtime)
	return runtime
}