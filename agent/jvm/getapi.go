package jvm

import (
	"monitor/agent/base"
	"fmt"
	//"monitor/agent/sock"
	"time"
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
	for _,arg := range args{
		for _,url := range baseUrl {
			_,res := base.HttpGet(url+method+"/java.lang:type="+arg)
			fmt.Println(string(res))
		}
	}
	time.Sleep(time.Duration(Frequency)*time.Second)
}