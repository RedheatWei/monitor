package jvm

import (
	"monitor/agent/base"
	"fmt"
)
//func getJson(baseurl,method,arg string) (n int,res []byte){
//	url := baseurl+method+"/java.lang:type="+arg
//	return base.HttpGet(url)
//}
func GetResJson(baseUrl []string,method,arg string){
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
//func GetRuntime(baseUrl []string){
//	for{
//		getResJson(baseUrl,"read","ClassLoading")
//		time.Sleep(5*time.Second)
//	}
//}



func AccceptGet(baseUrl []string,method,arg string){
	//go sock.ListenStart()
	//fmt.Println(string(getResJson(baseUrl,method,arg)))
	//return []byte{}
	//for{
	//	sock.SendMsg(string(js))
	//	time.Sleep(5*time.Second)
	//}

}