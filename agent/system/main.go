package system

import (
	"monitor/agent/base"
	"strconv"
)
var Frequency int64
var AgentConfig base.AgentConfig
//读取配置文件
func init()  {
	AgentConfig = base.GetConfig()
}
//启动
func Start() {
	Frequency,_ = strconv.ParseInt(AgentConfig.Default.Frequency,10,64)
	ch := make(chan int64)
	go CollectSysInfo()
	ch <- Frequency
}
