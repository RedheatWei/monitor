package jvm

import (
	"monitor/agent/base"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func getJson(port,method,arg string) (n int,res []byte){
	url := "http://127.0.0.1:"+port+"/jolokia/"+method+"/java.lang:type="+arg
	return base.HttpGet(url)
}
func GetRuntime(portlist []string){
	for _,port := range portlist{
		_,res := getJson(port,"read","Threading")
		js, js_err := simplejson.NewJson(res)
		fmt.Println(js_err)
		var nodes = make(map[string]interface{})
		nodes, _ = js.Map()
		for key,val := range nodes{
			fmt.Println(key,val)
		}
		//fmt.Println(string(res))
		//dec := json.NewDecoder(strings.NewReader(string(res)))
		//fmt.Println(dec)
	}
}