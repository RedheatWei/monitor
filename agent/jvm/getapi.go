package jvm

import (
	"monitor/base"
	//"fmt"
	"time"
	"monitor/network"
	"fmt"
)

func AccceptGet(baseUrl []string,args []string){
	for{
		for _,arg := range args{
			for _,url := range baseUrl {
				_,res := base.HttpGet(url+"read/java.lang:type="+arg)
				network.UdpSend(base.ReadAgentConfig("default","server"),res)
				fmt.Println(string(res))
			}
		}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}

func resGet(url,arg string) []byte{
	_,res := base.HttpGet(url+"read/java.lang:type="+arg)
	return  res
}