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
				network.UdpSend("123.56.92.243:33990",res)
				//hanJson(res)
			}
		}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}
func hanJson(res []byte){
	js,js_err := simplejson.NewJson(res)
	if js_err == nil {
		var nodes= make(map[string]interface{})
		nodes, _ = js.Map()
		//network.UdpSend("123.56.92.243:33990",nodes)
		fmt.Println(nodes)
	}

}