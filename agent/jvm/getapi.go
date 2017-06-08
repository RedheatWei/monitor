package jvm
import (
	"time"
	//"fmt"
	"monitor/agent/jvm/jvminfo"
	"encoding/json"
	"monitor/network"
	"monitor/base"
	"fmt"
)

type JsonInfo struct {
	//class
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit int64 `json:"HeapMemoryUsageInit"`
	HeapMemoryUsageCommitted int64 `json:"HeapMemoryUsageCommitted"`
	HeapMemoryUsageMax int64 `json:"HeapMemoryUsageMax"`
	HeapMemoryUsageUsed int64 `json:"HeapMemoryUsageUsed"`

	NonHeapMemoryUsageInit int64 `json:"NonHeapMemoryUsageInit"`
	NonHeapMemoryUsageCommitted int64 `json:"NonHeapMemoryUsageCommitted"`
	NonHeapMemoryUsageMax int64 `json:"NonHeapMemoryUsageMax"`
	NonHeapMemoryUsageUsed int64 `json:"NonHeapMemoryUsageUsed"`
	//rumtime
	ClassPath string `json:"ClassPath"`
	Uptime int64 `json:"Uptime"`
	//thread
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime int64 `json:"CurrentThreadCpuTime"`
	ThreadCount int32 `json:"ThreadCount"`
	DaemonThreadCount int32 `json:"DaemonThreadCount"`
	//base
	AgentId string `json:"agentId"`
	TimeStamp int64 `json:"TimeStamp"`
}

func dataHandle(url string) []byte{
	_,info := jvminfo.GetInfo(url)
	//if n != 1{
	//	return
	//}
	memory := jvminfo.GetMemory(url)
	runtime := jvminfo.GetRuntime(url)
	threading := jvminfo.GetThreading(url)
	classLoading := jvminfo.GetClassLoading(url)
	var allinfo JsonInfo
	//base
	allinfo.AgentId = info.Value.Config.AgentId
	allinfo.TimeStamp = info.TimeStamp
	//class
	allinfo.LoadedClassCount = classLoading.Value.LoadedClassCount
	allinfo.UnloadedClassCount = classLoading.Value.UnloadedClassCount
	allinfo.TotalLoadedClassCount = classLoading.Value.TotalLoadedClassCount
	//memory
	allinfo.HeapMemoryUsageInit = memory.Value.HeapMemoryUsage.Init
	allinfo.HeapMemoryUsageCommitted = memory.Value.HeapMemoryUsage.Committed
	allinfo.HeapMemoryUsageMax = memory.Value.HeapMemoryUsage.Max
	allinfo.HeapMemoryUsageUsed = memory.Value.HeapMemoryUsage.Used

	allinfo.NonHeapMemoryUsageInit = memory.Value.NonHeapMemoryUsage.Init
	allinfo.NonHeapMemoryUsageCommitted = memory.Value.NonHeapMemoryUsage.Committed
	allinfo.NonHeapMemoryUsageMax = memory.Value.NonHeapMemoryUsage.Max
	allinfo.NonHeapMemoryUsageUsed = memory.Value.NonHeapMemoryUsage.Used
	//runtime
	allinfo.ClassPath = runtime.Value.ClassPath
	allinfo.Uptime = runtime.Value.Uptime
	//thread
	allinfo.TotalStartedThreadCount = threading.Value.TotalStartedThreadCount
	allinfo.PeakThreadCount = threading.Value.PeakThreadCount
	allinfo.CurrentThreadCpuTime = threading.Value.CurrentThreadCpuTime
	allinfo.ThreadCount = threading.Value.ThreadCount
	allinfo.DaemonThreadCount = threading.Value.DaemonThreadCount
	fmt.Println(allinfo)
	i,err := json.Marshal(allinfo)
	fmt.Println(string(i))
	fmt.Println(err)
	return i
}

func AccceptGet(baseUrl []string,args []string){
	for{
		//for _,arg := range args{
			for _,url := range baseUrl {
				data := dataHandle(url)
				network.UdpSend(base.ReadAgentConfig("default","server"),data)
			}
		//}
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}