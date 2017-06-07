package jvm

import (
	"monitor/base"
	//"fmt"
	"time"
	"monitor/network"
	"fmt"
	"encoding/json"
)
type Memory struct {
	Request request `json:"request"`
	Value value `json:"value"`
	TimeStamp int32 `json:"timestamp"`
	Status int32 `json:"status"`
}
type memoryusage struct {
	Init int64 `json:"init"`
	Committed int64 `json:"committed"`
	Max int64 `json:"max"`
	Used int64 `json:"used"`
}
type value struct {
	ObjectPendingFinalizationCount int `json:"ObjectPendingFinalizationCount"`
	Verbose bool `json:"Verbose"`
	HeapMemoryUsage memoryusage `json:"HeapMemoryUsage"`
	NonHeapMemoryUsage memoryusage `json:"NonHeapMemoryUsage"`
	ObjectName memoryusage `json:"ObjectName"`
}
type request struct {
	Mbean string `json:"mbean"`
	Type string `json:"type"`
}
type objectname struct {
	objectName string `json:"objectName"`
}
func AccceptGet(baseUrl []string,method string,args []string){
	for{
		for _,arg := range args{
			for _,url := range baseUrl {
				_,res := base.HttpGet(url+method+"/java.lang:type="+arg)
				//fmt.Println(string(res))
				network.UdpSend(base.ReadAgentConfig("default","server"),res)
				//hanJson(res)
				//resJson := string(res)
				var memoryUsage Memory
				fmt.Println(json.Unmarshal(res,&memoryUsage))
				fmt.Println(memoryUsage)
			}
		}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}
//func hanJson(res []byte) []byte{
//	js,_ := simplejson.NewJson(res)
//	var nodes= make(map[string]interface{})
//	nodes, _ = js.Map()
//	return nodes["value"]
//}