package jvm

import (
	"monitor/agent/base"
	"fmt"
	//"encoding/json"
	//"strings"
)

func getJson(port,method,arg string) (n int,res []byte){
	url := "http://127.0.0.1:"+port+"/jolokia/"+method+"/java.lang:type="+arg
	return base.HttpGet(url)
}
func GetRuntime(portlist []string){
	for _,port := range portlist{
		_,res := getJson(port,"read","Threading")
		fmt.Println(string(res))
		//dec := json.NewDecoder(strings.NewReader(string(res)))
		//fmt.Println(dec)
	}
}