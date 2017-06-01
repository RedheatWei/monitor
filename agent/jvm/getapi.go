package jvm

import (
	"monitor/agent/base"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func getJson(baseurl,method,arg string) (n int,res []byte){
	url := baseurl+method+"/java.lang:type="+arg
	return base.HttpGet(url)
}
func getResJson(baseUrl []string,method,arg string){
	for _,url := range baseUrl{
		//_,res := getJson(url,"read","Threading")
		_,res := base.HttpGet(url+method+"/java.lang:type="+arg)
		js, js_err := simplejson.NewJson(res)
		if js_err == nil {
			var nodes= make(map[string]interface{})
			nodes, _ = js.Map()
			fmt.Println(nodes)
		}
	}
}
func GetRuntime(baseUrl []string){
	getResJson(baseUrl,"read","ClassLoading")
}