package system

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	//"github.com/bitly/go-simplejson"
	//"github.com/shirou/gopsutil/process"
	"monitor/base"
	"monitor/network"
	"time"
	"fmt"
	"encoding/json"
)

func CollectSysMemInfo() base.SysMemInfo{
	var SysMemInfo base.SysMemInfo
	v,_ := mem.VirtualMemory()
	s,_ := mem.SwapMemory()
	SysMemInfo.SwapMemoryStat = *s
	SysMemInfo.VirtualMemoryStat = *v
	return SysMemInfo
}
func CollectSysCpuInfo() base.SysCpuInfo{
	var SysCpuInfo base.SysCpuInfo
	i,_ := cpu.Info()
	//t,_ := cpu.Times()
	SysCpuInfo.InfoStat = i
	//SysCpuInfo.TimesStat = t
	return SysCpuInfo
}
func CollectSysDiskInfo() base.SysDiskInfo{
	var SysDiskInfo base.SysDiskInfo
	i,_ := disk.IOCounters()
	p,_ := disk.Partitions(false)
	//u,_ := disk.Usage()
	SysDiskInfo.IOCountersStat = i
	SysDiskInfo.PartitionStat = p
	return SysDiskInfo
}
func CollectSysHostInfo() base.SysHostInfo {
	var SysHostInfo base.SysHostInfo
	i,_ := host.Info()
	u,_ := host.Users()
	t,_ := host.SensorsTemperatures()
	SysHostInfo.InfoStat = *i
	SysHostInfo.UserStat = u
	SysHostInfo.TemperatureStat = t
	return SysHostInfo
}
func CollectSysLoadInfo() base.SysLoadInfo{
	var SysLoadInfo base.SysLoadInfo
	a,_ := load.Avg()
	m,_ := load.Misc()
	SysLoadInfo.AvgStat = *a
	SysLoadInfo.MiscStat = *m
	return SysLoadInfo
}
func CollectSysNetInfo() base.SysNetInfo{
	var SysNetInfo base.SysNetInfo
	i,_ := net.IOCounters(false)
	inf,_ := net.Interfaces()
	SysNetInfo.IOCountersStat = i
	SysNetInfo.InterfaceStat = inf
	return SysNetInfo
}
func CollectSysProcessInfo(){
	//var SysProcessInfo base.SysProcessInfo
}
func CollectSysInfo(){
	sysinfo := []string{"SysMemInfo","SysCpuInfo","SysDiskInfo","SysHostInfo","SysLoadInfo","SysNetInfo"}
	for _,t := range sysinfo{
		var data base.Senddata
		data.Type = t
		switch t {
		case "SysMemInfo":
			//data.SysMemInfo = CollectSysMemInfo()
			data.Data = CollectSysMemInfo()
		case "SysCpuInfo":
			//data.SysCpuInfo = CollectSysCpuInfo()
			data.Data = CollectSysCpuInfo()
		case "SysDiskInfo":
			//data.SysDiskInfo = CollectSysDiskInfo()
			data.Data = CollectSysDiskInfo()
		case "SysHostInfo":
			//data.SysHostInfo = CollectSysHostInfo()
			data.Data = CollectSysHostInfo()
		case "SysLoadInfo":
			//data.SysLoadInfo = CollectSysLoadInfo()
			data.Data = CollectSysLoadInfo()
		case "SysNetInfo":
			//data.SysNetInfo = CollectSysNetInfo()
			data.Data = CollectSysNetInfo()
		}
		fmt.Println(data)
		i,_ := json.Marshal(data)
		fmt.Println(i)
		network.UdpSend(base.ReadAgentConfig("default","server"),i)
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}
