package base

import (
	conf "monitor/server/config"
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
//配置信息
type ServerConfig struct {
	Default Default
	DB DB
}
type Default struct {
	Port string
}
type DB struct {
	Type string
	User string
	Passwd string
	Database string
	Host string
	Charset string
	Protocol string
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
type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}
type InterfaceAddr struct {
	Addr string `json:"addr"`
}
//读取服务端配置文件
func readServerConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/server.conf")
	return config.Read(mod,par)
}
func GetConfig() ServerConfig{
	var ServerConfig ServerConfig
	ServerConfig.Default.Port=readServerConfig("default","port")
	ServerConfig.DB.Type=readServerConfig("db","type")
	ServerConfig.DB.User=readServerConfig("db","user")
	ServerConfig.DB.Passwd=readServerConfig("db","passwd")
	ServerConfig.DB.Database=readServerConfig("db","database")
	ServerConfig.DB.Host=readServerConfig("db","host")
	ServerConfig.DB.Charset=readServerConfig("db","charset")
	ServerConfig.DB.Protocol=readServerConfig("db","protocol")
	return ServerConfig
}
