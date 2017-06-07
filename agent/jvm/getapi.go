package jvm

import (
	"monitor/base"
	"fmt"
	"time"
	"monitor/network"
	"encoding/json"
)
type Memory struct {
	Request request `json:"request"`
	Value value `json:"value"`
	TimeStamp int32 `json:"timestamp"`
	Status int32 `json:"status"`
}
type memoryusage struct {
	init int64 `json:"init"`
	committed int64 `json:"committed"`
	max int64 `json:"max"`
	used int64 `json:"used"`
}
type value struct {
	objectPendingFinalizationCount int `json:"ObjectPendingFinalizationCount"`
	verbose bool `json:"Verbose"`
	heapMemoryUsage interface{} `json:"HeapMemoryUsage"`
	nonHeapMemoryUsage interface{} `json:"NonHeapMemoryUsage"`
	objectName interface{} `json:"ObjectName"`
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
				resJson := string(res)
				var memoryUsage Memory
				fmt.Println(json.Unmarshal([]byte(resJson),&memoryUsage))
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