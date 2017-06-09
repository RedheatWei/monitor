package jvm
import (
	"time"
	//"fmt"
	"monitor/agent/jvm/jvminfo"
	"encoding/json"
	"monitor/network"
	"monitor/base"
)


func dataHandle(url string) []byte{
	_,info := jvminfo.GetInfo(url)
	//if n != 1{
	//	return
	//}
	memory := jvminfo.GetMemory(url)
	runtime := jvminfo.GetRuntime(url)
	threading := jvminfo.GetThreading(url)
	classLoading := jvminfo.GetClassLoading(url)
	var jvminfo base.JvmInfo
	//base
	jvminfo.AgentId = info.Value.Config.AgentId
	jvminfo.TimeStamp = info.TimeStamp
	//class
	jvminfo.LoadedClassCount = classLoading.Value.LoadedClassCount
	jvminfo.UnloadedClassCount = classLoading.Value.UnloadedClassCount
	jvminfo.TotalLoadedClassCount = classLoading.Value.TotalLoadedClassCount
	//memory
	jvminfo.HeapMemoryUsageInit = memory.Value.HeapMemoryUsage.Init
	jvminfo.HeapMemoryUsageCommitted = memory.Value.HeapMemoryUsage.Committed
	jvminfo.HeapMemoryUsageMax = memory.Value.HeapMemoryUsage.Max
	jvminfo.HeapMemoryUsageUsed = memory.Value.HeapMemoryUsage.Used

	jvminfo.NonHeapMemoryUsageInit = memory.Value.NonHeapMemoryUsage.Init
	jvminfo.NonHeapMemoryUsageCommitted = memory.Value.NonHeapMemoryUsage.Committed
	jvminfo.NonHeapMemoryUsageMax = memory.Value.NonHeapMemoryUsage.Max
	jvminfo.NonHeapMemoryUsageUsed = memory.Value.NonHeapMemoryUsage.Used
	//runtime
	jvminfo.ClassPath = runtime.Value.ClassPath
	jvminfo.Uptime = runtime.Value.Uptime
	//thread
	jvminfo.TotalStartedThreadCount = threading.Value.TotalStartedThreadCount
	jvminfo.PeakThreadCount = threading.Value.PeakThreadCount
	jvminfo.CurrentThreadCpuTime = threading.Value.CurrentThreadCpuTime
	jvminfo.ThreadCount = threading.Value.ThreadCount
	jvminfo.DaemonThreadCount = threading.Value.DaemonThreadCount
	var data base.Senddata
	data.Type = "jvminfo"
	data.JvmData = jvminfo
	i,_ := json.Marshal(data)
	return i
}
func CollectJvmInfo(baseUrl []string,args []string){
	for{
		for _,url := range baseUrl {
			data := dataHandle(url)
			network.UdpSend(base.ReadAgentConfig("default","server"),data)
		}
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}