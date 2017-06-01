package jvm

import (
	"monitor/agent/base"
)

func getJson(port,method,arg string) (n int,res []byte){
	url := "http://127.0.0.1:"+port+"/jolokia/"+method+"/"+arg
	return base.HttpGet(url)
}
