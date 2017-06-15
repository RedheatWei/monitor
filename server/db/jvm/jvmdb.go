package jvm

import (
	"monitor/server/base"
	"monitor/server/db"
	"fmt"
)

func InsertJvmDB(js base.JvmInfo,addr string){
	db := db.ConnDB()
	jvm := new(Collect_jvm)
	jvm.LoadedClassCount = js.LoadedClassCount
	jvm.UnloadedClassCount = js.UnloadedClassCount
	jvm.TotalLoadedClassCount = js.TotalLoadedClassCount
	jvm.HeapMemoryUsageInit = js.HeapMemoryUsageInit
	jvm.HeapMemoryUsageCommitted = js.HeapMemoryUsageCommitted
	jvm.HeapMemoryUsageMax = js.HeapMemoryUsageMax
	jvm.HeapMemoryUsageUsed = js.HeapMemoryUsageUsed
	jvm.NonHeapMemoryUsageInit = js.NonHeapMemoryUsageInit
	jvm.NonHeapMemoryUsageCommitted = js.NonHeapMemoryUsageCommitted
	jvm.NonHeapMemoryUsageMax = js.NonHeapMemoryUsageMax
	jvm.NonHeapMemoryUsageUsed = js.NonHeapMemoryUsageUsed
	jvm.ClassPath = js.ClassPath
	jvm.Uptime = js.Uptime
	jvm.TotalStartedThreadCount = js.TotalStartedThreadCount
	jvm.PeakThreadCount = js.PeakThreadCount
	jvm.CurrentThreadCpuTime = js.CurrentThreadCpuTime
	jvm.ThreadCount = js.ThreadCount
	jvm.DaemonThreadCount = js.DaemonThreadCount
	jvm.AgentId = js.AgentId
	jvm.TimeStamp = js.TimeStamp
	jvm.ServerIp = addr
	affected, err := db.Insert(jvm)
	fmt.Println(affected)
	fmt.Println(err)
	//db := ConnDB()
	//defer db.Close()
	//stmt, err := db.Prepare("INSERT INTO collect_jvm(LoadedClassCount,UnloadedClassCount,TotalLoadedClassCount,HeapMemoryUsageInit,HeapMemoryUsageCommitted,HeapMemoryUsageMax,HeapMemoryUsageUsed,NonHeapMemoryUsageInit,NonHeapMemoryUsageCommitted,NonHeapMemoryUsageMax,NonHeapMemoryUsageUsed,ClassPath,Uptime,TotalStartedThreadCount,PeakThreadCount,CurrentThreadCpuTime,ThreadCount,DaemonThreadCount,AgentId,TimeStamp,ServerIp) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//res, err := stmt.Exec(js.LoadedClassCount,js.UnloadedClassCount,js.TotalLoadedClassCount,js.HeapMemoryUsageInit,js.HeapMemoryUsageCommitted,js.HeapMemoryUsageMax,js.HeapMemoryUsageUsed,js.NonHeapMemoryUsageInit,js.NonHeapMemoryUsageCommitted,js.NonHeapMemoryUsageMax,js.NonHeapMemoryUsageUsed,js.ClassPath,js.Uptime,js.TotalStartedThreadCount,js.PeakThreadCount,js.CurrentThreadCpuTime,js.ThreadCount,js.DaemonThreadCount,js.AgentId,js.TimeStamp,addr)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//id,err := res.LastInsertId()
	//if err!=  nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(id)
}
