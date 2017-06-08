package jvm
import (
	"time"
	//"fmt"
	"monitor/agent/jvm/jvminfo"
	"encoding/json"
	"monitor/network"
	"monitor/base"
)
type jsoninfo struct {
	//class
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit int64 `json:"init"`
	HeapMemoryUsageCommitted int64 `json:"committed"`
	HeapMemoryUsageMax int64 `json:"max"`
	HeapMemoryUsageUsed int64 `json:"used"`
	NonHeapMemoryUsageInit int64 `json:"init"`
	NonHeapMemoryUsageCommitted int64 `json:"committed"`
	NonHeapMemoryUsageMax int64 `json:"max"`
	NonHeapMemoryUsageUsed int64 `json:"used"`
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
	info := jvminfo.GetInfo(url)
	memory := jvminfo.GetMemory(url)
	runtime := jvminfo.GetRuntime(url)
	threading := jvminfo.GetThreading(url)
	classLoading := jvminfo.GetClassLoading(url)
	var allinfo jsoninfo
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
	i,_ := json.Marshal(allinfo)
	return i
}

func AccceptGet(baseUrl []string,args []string){
	for{
		//for _,arg := range args{
			for _,url := range baseUrl {
				data := dataHandle(url)
				//fmt.Println(string(data))
				//fmt.Println(jvminfo.GetMemory(url).Value.HeapMemoryUsage.Committed)
				//_,res := base.HttpGet(url+"read/java.lang:type="+arg)
				network.UdpSend(base.ReadAgentConfig("default","server"),data)
			}
		//}
		time.Sleep(time.Duration(Frequency)*time.Second)
	}
}