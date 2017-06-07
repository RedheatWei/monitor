package jvm

import (
	"monitor/base"
	"fmt"
	//"monitor/agent/sock"
	"time"
	"github.com/bitly/go-simplejson"
	"monitor/network"
)
func getResJson(baseUrl []string,method,arg string){
	for _,url := range baseUrl{
		//_,res := getJson(url,"read","Threading")
		_,res := base.HttpGet(url+method+"/java.lang:type="+arg)
		//js, js_err := simplejson.NewJson(res)
		fmt.Println(string(res))
		//return res
		//if js_err == nil {
		//	var nodes= make(map[string]interface{})
		//	nodes, _ = js.Map()
		//	return nodes
		//}
	}
	//return []byte{}
}
func AccceptGet(baseUrl []string,method string,args []string){
	for{
		for _,arg := range args{
			for _,url := range baseUrl {
				_,res := base.HttpGet(url+method+"/java.lang:type="+arg)
				//fmt.Println(string(res))
				network.UdpSend(base.ReadAgentConfig("default","server"),res)
				hanJson(res)
			}
		}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}
func hanJson(res []byte){
	js,_ := simplejson.NewJson(res)
	var nodes= make(map[string]interface{})
	nodes, _ = js.Map()
	//fmt.Println(nodes["value"])
	js2,_ := simplejson.NewJson(nodes["value"].([]byte))
	var nodes2 = make(map[string]interface{})
	nodes2,_ = js2.Map()
	fmt.Println(nodes2["HeapMemoryUsage"])

}