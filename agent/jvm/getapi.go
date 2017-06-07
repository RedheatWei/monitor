package jvm

import (
	"monitor/base"
	//"fmt"
	"time"
	"monitor/network"
	"fmt"
	"encoding/json"
)

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