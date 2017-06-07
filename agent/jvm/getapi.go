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
	init string
	committed string
	max string
	used string
}
type value struct {
	ObjectPendingFinalizationCount string
	Verbose string
	HeapMemoryUsage memoryusage
	NonHeapMemoryUsage memoryusage
	ObjectName string
	TimeStamp string
	Status string
}
type request struct {
	Mbean string `json:"mbean"`
	Type string `json:"type"`
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
				fmt.Println(memoryUsage.Value.ObjectPendingFinalizationCount)

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