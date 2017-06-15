package system

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	//"github.com/shirou/gopsutil/process"
	"monitor/agent/base"
	"monitor/agent/network"
	"time"
	"encoding/json"
)
var Server string
//读取配置文件
func init()  {
	Server = AgentConfig.Default.Server
}
//收集系统内存信息
func CollectSysMemInfo(){
	var SysMemInfo base.SysMemInfo
	v,_ := mem.VirtualMemory()
	s,_ := mem.SwapMemory()
	SysMemInfo.SwapMemoryStat = *s
	SysMemInfo.VirtualMemoryStat = *v
	SysMemInfo.Type = "SysMemInfo"
	by,_ := json.Marshal(SysMemInfo)
	network.UdpSend(Server,by)
}
//收集系统cpu信息
func CollectSysCpuInfo(){
	var SysCpuInfo base.SysCpuInfo
	i,_ := cpu.Info()
	SysCpuInfo.InfoStat = i
	SysCpuInfo.Type = "SysCpuInfo"
	by,_ := json.Marshal(SysCpuInfo)
	network.UdpSend(Server,by)
}
//收集系统磁盘信息
func CollectSysDiskInfo(){
	var SysDiskInfo base.SysDiskInfo
	i,_ := disk.IOCounters()
	p,_ := disk.Partitions(false)
	SysDiskInfo.IOCountersStat = i
	SysDiskInfo.PartitionStat = p
	SysDiskInfo.Type = "SysDiskInfo"
	by,_ := json.Marshal(SysDiskInfo)
	network.UdpSend(Server,by)
}
//收集系统主机信息
func CollectSysHostInfo() {
	var SysHostInfo base.SysHostInfo
	i,_ := host.Info()
	u,_ := host.Users()
	t,_ := host.SensorsTemperatures()
	SysHostInfo.InfoStat = *i
	SysHostInfo.UserStat = u
	SysHostInfo.TemperatureStat = t
	SysHostInfo.Type = "SysHostInfo"
	by,_ := json.Marshal(SysHostInfo)
	network.UdpSend(Server,by)
}
//收集系统负载信息
func CollectSysLoadInfo() {
	var SysLoadInfo base.SysLoadInfo
	a,_ := load.Avg()
	m,_ := load.Misc()
	SysLoadInfo.AvgStat = *a
	SysLoadInfo.MiscStat = *m
	SysLoadInfo.Type = "SysLoadInfo"
	by,_ := json.Marshal(SysLoadInfo)
	network.UdpSend(Server,by)
}
//收集系统网络信息
func CollectSysNetInfo() {
	var SysNetInfo base.SysNetInfo
	i,_ := net.IOCounters(false)
	inf,_ := net.Interfaces()
	SysNetInfo.IOCountersStat = i
	SysNetInfo.InterfaceStat = inf
	SysNetInfo.Type = "SysNetInfo"
	by,_ := json.Marshal(SysNetInfo)
	network.UdpSend(Server,by)
}
//收集系统进程信息,没在使用
func CollectSysProcessInfo(){
	//var SysProcessInfo base.SysProcessInfo
}
//收集信息并发送
func CollectSysInfo(){
	sysinfo := []string{"SysMemInfo","SysCpuInfo","SysDiskInfo","SysHostInfo","SysLoadInfo","SysNetInfo"}
	for _,t := range sysinfo{
		switch t {
		case "SysMemInfo":
			CollectSysMemInfo()
		case "SysCpuInfo":
			CollectSysCpuInfo()
		case "SysDiskInfo":
			CollectSysDiskInfo()
		case "SysHostInfo":
			CollectSysHostInfo()
		case "SysLoadInfo":
			CollectSysLoadInfo()
		case "SysNetInfo":
			CollectSysNetInfo()
		}
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}
