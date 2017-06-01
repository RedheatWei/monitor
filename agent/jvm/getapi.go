package jvm

import (
	"monitor/agent/base"
	"fmt"
	//"encoding/json"
	//"strings"
)

func getJson(port,method,arg string) (n int,res []byte){
	url := "http://127.0.0.1:"+port+"/jolokia/"+method+"/"+arg
	return base.HttpGet(url)
}
func GetRuntime(portlist []string){
	for _,port := range portlist{
		_,res := getJson(port,"read","java.lang:type=Runtime")
		fmt.Println(string(res))
		//dec := json.NewDecoder(strings.NewReader(string(res)))
		//fmt.Println(dec)
	}
}