package jvm

import (
	"monitor/agent/base"
	"fmt"
)

func getJson(port,method,arg string) (n int,res []byte){
	url := "http://127.0.0.1:"+port+"/jolokia/"+method+"/"+arg
	return base.HttpGet(url)
}
func GetRuntime(portlist []string){
	for _,port := range portlist{
		_,res := getJson(port,"read","java.lang:type=Runtime")
		fmt.Println(string(res))
	}
}