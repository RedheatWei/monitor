package system

import (
	"monitor/agent/base"
	"strconv"
)
var Frequency int64
var AgentConfig base.AgentConfig
//启动
func Start() {
	AgentConfig = base.GetConfig()
	Frequency,_ = strconv.ParseInt(AgentConfig.Default.Frequency,10,64)
	CollectSysInfo()
}
