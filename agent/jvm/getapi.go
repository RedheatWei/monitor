package jvm

import (
	"time"
	"fmt"
	"monitor/agent/jvm/jvminfo"
)

func AccceptGet(baseUrl []string,args []string){
	for{
		//for _,arg := range args{
			for _,url := range baseUrl {
				fmt.Println(jvminfo.GetMemory(url).Value.HeapMemoryUsage.Committed)
				//_,res := base.HttpGet(url+"read/java.lang:type="+arg)
				//network.UdpSend(base.ReadAgentConfig("default","server"),res)
			}
		//}
		time.Sleep(time.Duration(Frequency)*time.Second)

	}
}

