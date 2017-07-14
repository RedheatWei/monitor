package jvm

import (
	"monitor/server/base"
	"fmt"
	"time"
	"reflect"
	"encoding/json"
	"monitor/server/db/opentsdb"
)
//定义收集的项目
type jvmDB struct {
	LoadedClassCount *opentsdb.Collect_jvm
	UnloadedClassCount *opentsdb.Collect_jvm
	TotalLoadedClassCount *opentsdb.Collect_jvm
	HeapMemoryUsageCommitted *opentsdb.Collect_jvm
	HeapMemoryUsageUsed *opentsdb.Collect_jvm
	NonHeapMemoryUsageCommitted *opentsdb.Collect_jvm
	NonHeapMemoryUsageUsed *opentsdb.Collect_jvm
	TotalStartedThreadCount *opentsdb.Collect_jvm
	PeakThreadCount *opentsdb.Collect_jvm
	ThreadCount *opentsdb.Collect_jvm
	DaemonThreadCount *opentsdb.Collect_jvm
	HeapMemoryUsageInit *opentsdb.Collect_jvm
	HeapMemoryUsageMax *opentsdb.Collect_jvm
	NonHeapMemoryUsageInit *opentsdb.Collect_jvm
	NonHeapMemoryUsageMax *opentsdb.Collect_jvm
}
//主函数
func InsertJvmDB(js base.JvmInfo,serverIpInfo base.ServerIpInfo){
	var jvmDB jvmDB
	jvmDB.LoadedClassCount = loadedClassCount(js,serverIpInfo)
	jvmDB.UnloadedClassCount = unloadedClassCount(js,serverIpInfo)
	jvmDB.TotalLoadedClassCount = totalLoadedClassCount(js,serverIpInfo)
	jvmDB.HeapMemoryUsageCommitted = heapMemoryUsageCommitted(js,serverIpInfo)
	jvmDB.HeapMemoryUsageUsed = heapMemoryUsageUsed(js,serverIpInfo)
	jvmDB.NonHeapMemoryUsageCommitted = nonHeapMemoryUsageCommitted(js,serverIpInfo)
	jvmDB.NonHeapMemoryUsageUsed = nonHeapMemoryUsageUsed(js,serverIpInfo)
	jvmDB.TotalStartedThreadCount = totalStartedThreadCount(js,serverIpInfo)
	jvmDB.PeakThreadCount = peakThreadCount(js,serverIpInfo)
	jvmDB.ThreadCount = threadCount(js,serverIpInfo)
	jvmDB.DaemonThreadCount = daemonThreadCount(js,serverIpInfo)
	jvmDB.HeapMemoryUsageInit = heapMemoryUsageInit(js,serverIpInfo)
	jvmDB.HeapMemoryUsageMax = heapMemoryUsageMax(js,serverIpInfo)
	jvmDB.NonHeapMemoryUsageInit = nonHeapMemoryUsageInit(js,serverIpInfo)
	jvmDB.NonHeapMemoryUsageMax = nonHeapMemoryUsageMax(js,serverIpInfo)
	v := reflect.ValueOf(jvmDB)
	var sli_str []interface{}
	for k := 0; k < v.NumField(); k++{
		if ! v.Field(k).IsNil(){
			val := v.Field(k).Interface()
			sli_str = append(sli_str,val)
		}
	}
	b,err := json.Marshal(sli_str)
	if err!=nil{
		fmt.Println(err)
	}
	opentsdb.SendToTsDb(string(b))
}
//组合数据,便于修改tag
func createCollect(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := new(opentsdb.Collect_jvm)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = serverIpInfo.Server
	collect.Tags.Ip = serverIpInfo.Ip
	collect.Tags.CtimeStamp = js.TimeStamp
	collect.Tags.ClassPath = js.ClassPath
	collect.Tags.AgentId = js.AgentId
	collect.Tags.CurrentThreadCpuTime = js.CurrentThreadCpuTime
	collect.Tags.Uptime = js.Uptime
	return collect
}
func loadedClassCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.loadedClassCount"
	collect.Value = float64(js.LoadedClassCount)
	return collect
}
func unloadedClassCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.unloadedClassCount"
	collect.Value = float64(js.UnloadedClassCount)
	return collect
}
func totalLoadedClassCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.totalLoadedClassCount"
	collect.Value = float64(js.TotalLoadedClassCount)
	return collect
}
func heapMemoryUsageCommitted(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.heapMemoryUsageCommitted"
	collect.Value = float64(js.HeapMemoryUsageCommitted)
	return collect
}
func heapMemoryUsageUsed(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.heapMemoryUsageUsed"
	collect.Value = float64(js.HeapMemoryUsageUsed)
	return collect
}
func nonHeapMemoryUsageCommitted(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.nonHeapMemoryUsageCommitted"
	collect.Value = float64(js.NonHeapMemoryUsageCommitted)
	return collect
}
func nonHeapMemoryUsageUsed(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.nonHeapMemoryUsageUsed"
	collect.Value = float64(js.NonHeapMemoryUsageUsed)
	return collect
}
func totalStartedThreadCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.totalStartedThreadCount"
	collect.Value = float64(js.TotalStartedThreadCount)
	return collect
}
func peakThreadCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.peakThreadCount"
	collect.Value = float64(js.PeakThreadCount)
	return collect
}
func threadCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.threadCount"
	collect.Value = float64(js.ThreadCount)
	return collect
}
func daemonThreadCount(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.daemonThreadCount"
	collect.Value = float64(js.DaemonThreadCount)
	return collect
}
func heapMemoryUsageInit(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.heapMemoryUsageInit"
	collect.Value = float64(js.HeapMemoryUsageInit)
	return collect
}
func heapMemoryUsageMax(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.heapMemoryUsageMax"
	collect.Value = float64(js.HeapMemoryUsageMax)
	return collect
}
func nonHeapMemoryUsageInit(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.nonHeapMemoryUsageInit"
	collect.Value = float64(js.NonHeapMemoryUsageInit)
	return collect
}
func nonHeapMemoryUsageMax(js base.JvmInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect_jvm{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.jvm.nonHeapMemoryUsageMax"
	collect.Value = float64(js.NonHeapMemoryUsageMax)
	return collect
}
