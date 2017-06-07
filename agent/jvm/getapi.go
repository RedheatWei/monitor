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
}
type memoryusage struct {
	init string `json:"init"`
	committed string `json:"committed"`
	max string `json:"max"`
	used string `json:"used"`
}
type value struct {
	objectpendingfinalizationcount string `json:"ObjectPendingFinalizationCount"`
	verbose string `json:"Verbose"`
	heapmemoryusage memoryusage `json:"HeapMemoryUsage"`
	nonheapmemoryusage memoryusage `json:"NonHeapMemoryUsage"`
	ObjectName objectname `json:"ObjectName"`
}
type request struct {
	Mbean string `json:"mbean"`
	Type string `json:"type"`
	TimeStamp string `json:"timestamp"`
	Status string `json:"status"`
}
type objectname struct {
	objectname string `json:"objectName"`
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
				json.Unmarshal([]byte(resJson),&memoryUsage)
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