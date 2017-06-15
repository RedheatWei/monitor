package base

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"monitor/conf"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	//"github.com/shirou/gopsutil/process"
)
type JvmInfo struct {
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
type SysMemInfo struct {
	VirtualMemoryStat mem.VirtualMemoryStat `json:"VirtualMemoryStat"`
	SwapMemoryStat mem.SwapMemoryStat `json:"SwapMemoryStat"`
}
type SysCpuInfo struct {
	InfoStat []cpu.InfoStat `json:"InfoStat"`
	//TimesStat []cpu.TimesStat `json:"TimesStat"`
}
type SysDiskInfo struct {
	IOCountersStat map[string]disk.IOCountersStat `json:"IOCountersStat"`
	PartitionStat []disk.PartitionStat `json:"PartitionStat"`
	//UsageStat disk.UsageStat `json:"UsageStat"`
}
type SysHostInfo struct {
	InfoStat host.InfoStat `json:"InfoStat"`
	TemperatureStat []host.TemperatureStat `json:"TemperatureStat"`
	UserStat []host.UserStat `json:"UserStat"`
}
type SysLoadInfo struct {
	AvgStat load.AvgStat `json:"AvgStat"`
	MiscStat load.MiscStat `json:"MiscStat"`
}
type SysNetInfo struct {
	//ConnectionStat net.ConnectionStat `json:"ConnectionStat"`
	IOCountersStat []net.IOCountersStat `json:"IOCountersStat"`
	InterfaceStat []net.InterfaceStat `json:"InterfaceStat"`
	//ProtoCountersStat net.ProtoCountersStat `json:"ProtoCountersStat"`
}
type SysProcessInfo struct {
	//IOCountersStat process.IOCountersStat `json:"IOCountersStat"`
	//MemoryInfoExStat process.MemoryInfoExStat `json:"MemoryInfoExStat"`
	//MemoryInfoStat process.MemoryInfoStat `json:"MemoryInfoStat"`
	//MemoryMapsStat process.MemoryMapsStat `json:"MemoryMapsStat"`
	//NumCtxSwitchesStat process.NumCtxSwitchesStat `json:"NumCtxSwitchesStat"`
	//OpenFilesStat process.OpenFilesStat `json:"OpenFilesStat"`
	//Process process.Process `json:"Process"`
	//RlimitStat process.RlimitStat `json:"RlimitStat"`
}
type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}
type InterfaceAddr struct {
	Addr string `json:"addr"`
}
type Senddata struct {
	Type           string `json:"Type"`
	Data interface{}
	//JvmInfo        JvmInfo `json:"JvmInfo"`
	//SysMemInfo     SysMemInfo `json:"SysMemInfo"`
	//SysCpuInfo     SysCpuInfo `json:"SysCpuInfo"`
	//SysDiskInfo    SysDiskInfo `json:"SysDiskInfo"`
	//SysHostInfo    SysHostInfo `json:"SysHostInfo"`
	//SysLoadInfo    SysLoadInfo `json:"SysLoadInfo"`
	//SysNetInfo     SysNetInfo `json:"SysNetInfo"`
	//SysProcessInfo SysProcessInfo `json:"SysProcessInfo"`
}

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
func ReadAgentConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/agent.conf")
	return config.Read(mod,par)
}
//读取服务端配置文件
func ReadServerConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/server.conf")
	return config.Read(mod,par)
}
