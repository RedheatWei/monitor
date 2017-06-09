package system

import (
	"monitor/base"
	"strconv"
)
var Frequency int64
//启动
func Start() {
	Frequency,_ = strconv.ParseInt(base.ReadAgentConfig("default","frequency"),10,64)
	//ch := make(chan []string)
	//go CollectInfo(baseUrl,args)
	//ch <- baseUrl
}
