package system

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	//"github.com/shirou/gopsutil/process"
	"monitor/base"
	"monitor/network"
	"time"
	"encoding/json"
)
var Server string
func init(){
	Server = base.ReadAgentConfig("default","server")
}

func CollectSysMemInfo(){
	var SysMemInfo base.SysMemInfo
	v,_ := mem.VirtualMemory()
	s,_ := mem.SwapMemory()
	SysMemInfo.SwapMemoryStat = *s
	SysMemInfo.VirtualMemoryStat = *v
	SysMemInfo.Type = "SysMemInfo"
	by,_ := json.Marshal(SysMemInfo)
	network.UdpSend(Server,by)
	//return SysMemInfo
}
func CollectSysCpuInfo(){
	var SysCpuInfo base.SysCpuInfo
	i,_ := cpu.Info()
	//t,_ := cpu.Times()
	SysCpuInfo.InfoStat = i
	SysCpuInfo.Type = "SysCpuInfo"
	//SysCpuInfo.TimesStat = t
	//return SysCpuInfo
	by,_ := json.Marshal(SysCpuInfo)
	network.UdpSend(Server,by)
}
func CollectSysDiskInfo(){
	var SysDiskInfo base.SysDiskInfo
	i,_ := disk.IOCounters()
	p,_ := disk.Partitions(false)
	//u,_ := disk.Usage()
	SysDiskInfo.IOCountersStat = i
	SysDiskInfo.PartitionStat = p
	SysDiskInfo.Type = "SysDiskInfo"
	//return SysDiskInfo
	by,_ := json.Marshal(SysDiskInfo)
	network.UdpSend(Server,by)
}
func CollectSysHostInfo() {
	var SysHostInfo base.SysHostInfo
	i,_ := host.Info()
	u,_ := host.Users()
	t,_ := host.SensorsTemperatures()
	SysHostInfo.InfoStat = *i
	SysHostInfo.UserStat = u
	SysHostInfo.TemperatureStat = t
	SysHostInfo.Type = "SysHostInfo"
	//return SysHostInfo
	by,_ := json.Marshal(SysHostInfo)
	network.UdpSend(Server,by)
}
func CollectSysLoadInfo() {
	var SysLoadInfo base.SysLoadInfo
	a,_ := load.Avg()
	m,_ := load.Misc()
	SysLoadInfo.AvgStat = *a
	SysLoadInfo.MiscStat = *m
	SysLoadInfo.Type = "SysLoadInfo"
	//return SysLoadInfo
	by,_ := json.Marshal(SysLoadInfo)
	network.UdpSend(Server,by)
}
func CollectSysNetInfo() {
	var SysNetInfo base.SysNetInfo
	i,_ := net.IOCounters(false)
	inf,_ := net.Interfaces()
	SysNetInfo.IOCountersStat = i
	SysNetInfo.InterfaceStat = inf
	SysNetInfo.Type = "SysNetInfo"
	by,_ := json.Marshal(SysNetInfo)
	network.UdpSend(Server,by)
	//return SysNetInfo
}
func CollectSysProcessInfo(){
	//var SysProcessInfo base.SysProcessInfo
}
func CollectSysInfo(){
	sysinfo := []string{"SysMemInfo","SysCpuInfo","SysDiskInfo","SysHostInfo","SysLoadInfo","SysNetInfo"}
	for _,t := range sysinfo{
		//var data base.Senddata
		//data.Type = t
		switch t {
		case "SysMemInfo":
			//data.SysMemInfo = CollectSysMemInfo()
			CollectSysMemInfo()
		case "SysCpuInfo":
			//data.SysCpuInfo = CollectSysCpuInfo()
			CollectSysCpuInfo()
		case "SysDiskInfo":
			//data.SysDiskInfo = CollectSysDiskInfo()
			CollectSysDiskInfo()
		case "SysHostInfo":
			//data.SysHostInfo = CollectSysHostInfo()
			CollectSysHostInfo()
		case "SysLoadInfo":
			//data.SysLoadInfo = CollectSysLoadInfo()
			CollectSysLoadInfo()
		case "SysNetInfo":
			//data.SysNetInfo = CollectSysNetInfo()
			CollectSysNetInfo()
		}
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}
