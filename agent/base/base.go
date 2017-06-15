package base

import (
	"net/http"
	"fmt"
	"io/ioutil"
	conf "monitor/agent/config"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	//"github.com/shirou/gopsutil/process"
)
//jvm信息
type JvmInfo struct {
	Type           string `json:"Type"`
	//class
	LoadedClassCount      int32 `json:"LoadedClassCount"`
	UnloadedClassCount    int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit      int64 `json:"HeapMemoryUsageInit"`
	HeapMemoryUsageCommitted int64 `json:"HeapMemoryUsageCommitted"`
	HeapMemoryUsageMax       int64 `json:"HeapMemoryUsageMax"`
	HeapMemoryUsageUsed      int64 `json:"HeapMemoryUsageUsed"`

	NonHeapMemoryUsageInit      int64 `json:"NonHeapMemoryUsageInit"`
	NonHeapMemoryUsageCommitted int64 `json:"NonHeapMemoryUsageCommitted"`
	NonHeapMemoryUsageMax       int64 `json:"NonHeapMemoryUsageMax"`
	NonHeapMemoryUsageUsed      int64 `json:"NonHeapMemoryUsageUsed"`
	//rumtime
	ClassPath string `json:"ClassPath"`
	Uptime    int64 `json:"Uptime"`
	//thread
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount         int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime    int64 `json:"CurrentThreadCpuTime"`
	ThreadCount             int32 `json:"ThreadCount"`
	DaemonThreadCount       int32 `json:"DaemonThreadCount"`
	//base
	AgentId   string `json:"agentId"`
	TimeStamp int64 `json:"TimeStamp"`
}
//客户端配置
type AgentConfig struct {
	Default Default
	Jvm Jvm
}
type Default struct {
	Server string
	Frequency string
}
type Jvm struct {
	Jolokiapath string
	Jolokianame string
	Portstart string
}
//系统内存信息
type SysMemInfo struct {
	Type           string `json:"Type"`
	VirtualMemoryStat mem.VirtualMemoryStat `json:"VirtualMemoryStat"`
	SwapMemoryStat mem.SwapMemoryStat `json:"SwapMemoryStat"`
}
//系统cpu信息
type SysCpuInfo struct {
	Type           string `json:"Type"`
	InfoStat []cpu.InfoStat `json:"InfoStat"`
	//TimesStat []cpu.TimesStat `json:"TimesStat"`
}
//系统磁盘信息
type SysDiskInfo struct {
	Type           string `json:"Type"`
	IOCountersStat map[string]disk.IOCountersStat `json:"IOCountersStat"`
	PartitionStat []disk.PartitionStat `json:"PartitionStat"`
	//UsageStat disk.UsageStat `json:"UsageStat"`
}
//系统主机信息
type SysHostInfo struct {
	Type           string `json:"Type"`
	InfoStat host.InfoStat `json:"InfoStat"`
	TemperatureStat []host.TemperatureStat `json:"TemperatureStat"`
	UserStat []host.UserStat `json:"UserStat"`
}
//系统负载信息
type SysLoadInfo struct {
	Type           string `json:"Type"`
	AvgStat load.AvgStat `json:"AvgStat"`
	MiscStat load.MiscStat `json:"MiscStat"`
}
//系统网络信息
type SysNetInfo struct {
	Type           string `json:"Type"`
	//ConnectionStat net.ConnectionStat `json:"ConnectionStat"`
	IOCountersStat []net.IOCountersStat `json:"IOCountersStat"`
	InterfaceStat []net.InterfaceStat `json:"InterfaceStat"`
	//ProtoCountersStat net.ProtoCountersStat `json:"ProtoCountersStat"`
}
//系统进程信息
type SysProcessInfo struct {
	Type           string `json:"Type"`
	//IOCountersStat process.IOCountersStat `json:"IOCountersStat"`
	//MemoryInfoExStat process.MemoryInfoExStat `json:"MemoryInfoExStat"`
	//MemoryInfoStat process.MemoryInfoStat `json:"MemoryInfoStat"`
	//MemoryMapsStat process.MemoryMapsStat `json:"MemoryMapsStat"`
	//NumCtxSwitchesStat process.NumCtxSwitchesStat `json:"NumCtxSwitchesStat"`
	//OpenFilesStat process.OpenFilesStat `json:"OpenFilesStat"`
	//Process process.Process `json:"Process"`
	//RlimitStat process.RlimitStat `json:"RlimitStat"`
}
//type Addr struct {
//	IP   string `json:"ip"`
//	Port uint32 `json:"port"`
//}
//type InterfaceAddr struct {
//	Addr string `json:"addr"`
//}
//type Senddata struct {
//	Data interface{}
//	//Type           string `json:"Type"`
//	//JvmInfo        JvmInfo `json:"JvmInfo"`
//	//SysMemInfo     SysMemInfo `json:"SysMemInfo"`
//	//SysCpuInfo     SysCpuInfo `json:"SysCpuInfo"`
//	//SysDiskInfo    SysDiskInfo `json:"SysDiskInfo"`
//	//SysHostInfo    SysHostInfo `json:"SysHostInfo"`
//	//SysLoadInfo    SysLoadInfo `json:"SysLoadInfo"`
//	//SysNetInfo     SysNetInfo `json:"SysNetInfo"`
//	//SysProcessInfo SysProcessInfo `json:"SysProcessInfo"`
//}

//get方法
//get方法
func HttpGet(url string) (n int,res []byte){
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1,[]byte{}
	}else{
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		return 0,body
	}
}
//读取配置文件
func readAgentConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/agent.conf")
	return config.Read(mod,par)
}
//读取配置文件
func GetConfig() AgentConfig{
	var AgentConfig AgentConfig
	AgentConfig.Default.Server=readAgentConfig("default","server")
	AgentConfig.Default.Frequency=readAgentConfig("default","frequency")
	AgentConfig.Jvm.Jolokiapath=readAgentConfig("jvm","jolokiapath")
	AgentConfig.Jvm.Jolokianame=readAgentConfig("jvm","jolokianame")
	AgentConfig.Jvm.Portstart=readAgentConfig("jvm","portstart")
	return AgentConfig
}
