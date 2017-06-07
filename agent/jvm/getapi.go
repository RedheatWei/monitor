package jvm

import (
	"monitor/base"
	//"fmt"
	"time"
	//"monitor/network"
	"fmt"
	"monitor/agent/jvm/jvminfo"
)

func AccceptGet(baseUrl []string,args []string){
	for{
		//for _,arg := range args{
			for _,url := range baseUrl {
				fmt.Println(jvminfo.GetMemory(url))
				//_,res := base.HttpGet(url+"read/java.lang:type="+arg)
				//network.UdpSend(base.ReadAgentConfig("default","server"),res)
			}
		//}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}

func ResGet(url,arg string) []byte{
	_,res := base.HttpGet(url+"read/java.lang:type="+arg)
	return  res
}